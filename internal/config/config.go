package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// LoadConfig loads configuration for a specific application.
// appName corresponds to the config file prefix (e.g., "api" for api.yaml)
func LoadConfig(app App, config any) error {
	v := viper.New()

	// 1. Set default values
	val := reflect.ValueOf(config)
	applyDefaultsRecursive(v, app, val)

	// 2. Set config file name and paths
	// Config file name will be e.g., "api.yaml", "worker.yaml"
	configFileName := app.Name
	if os.Getenv("APP_CONFIG_NAME") != "" {
		// Allow overriding via APP_CONFIG_NAME for specific app
		configFileName = os.Getenv("APP_CONFIG_NAME")
	}

	v.SetConfigName(configFileName)
	v.AddConfigPath("./configs")                     // Search in the `configs` directory
	v.AddConfigPath(".")                             // Search in current directory (for development convenience)
	v.AddConfigPath(filepath.Join("/etc", app.Name)) // Linux system config

	// 3. Read config file
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Printf("Warning: Config file '%s.yaml' not found, loading from environment variables and defaults.\n", configFileName)
		}

		// 4. Environment variables
		v.SetEnvPrefix(strings.ToUpper(app.Name)) // Each app can have its own env prefix (e.g., API_DB_HOST, WORKER_QUEUE_NAME)
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		v.AutomaticEnv()
	}

	// 5. Unmarshal and set app name
	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("failed to unmarshal config for %s: %w", app.Name, err)
	}

	return nil
}

// DefaultSetter 定义接口，任何 struct 实现它就能在 LoadConfig 被调用
type DefaultSetter interface {
	SetDefault(v *viper.Viper, app App)
}

// 递归遍历 struct，查找实现 DefaultSetter 的字段并调用其方法
func applyDefaultsRecursive(v *viper.Viper, app App, val reflect.Value) {
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			if val.CanSet() {
				val.Set(reflect.New(val.Type().Elem()))
			} else {
				// val 是嵌套字段，不能直接 set，需要在父 struct 那里设置（见下）
				return
			}
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	// 如果当前 struct 实现了 SetDefault，则调用
	if setter, ok := val.Addr().Interface().(DefaultSetter); ok {
		setter.SetDefault(v, app)
	}

	// 遍历所有字段，递归处理
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		applyDefaultsRecursive(v, app, field)
	}
}

type Common struct {
	Logger *Logger `mapstructure:"logger"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

func (s App) String() string {
	return fmt.Sprintf("%s[%s]", s.Name, s.Version)
}

type Logger struct {
	File  string `mapstructure:"file"`
	Level string `mapstructure:"level"`
	// MaxSize log maximum size, unit is mb
	MaxSize int `mapstructure:"maxsize"`
	// MaxBackups maximum log backup
	MaxBackups int `mapstructure:"maxBackups"`
	// MaxAge backup log maximum age, unit is day
	MaxAge   int  `mapstructure:"maxAge"`
	Compress bool `mapstructure:"compress"`
}

func (l *Logger) SetDefault(v *viper.Viper, service App) {
	v.SetDefault("logger.level", "info")
	v.SetDefault("logger.file", fmt.Sprintf("logs/%s.log", service.Name))
	v.SetDefault("logger.maxsize", 128)
	v.SetDefault("logger.maxBackups", 3)
	v.SetDefault("logger.maxAge", 7)
}
