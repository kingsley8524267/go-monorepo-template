package config

type MyApp struct {
	Common `mapstructure:",squash"`
}
