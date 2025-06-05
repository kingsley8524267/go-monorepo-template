package logger

import (
	"errors"
	"fmt"
	"go-monorepo-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *Logger

type Logger struct {
	config *config.Logger

	writer *zap.Logger
	sugar  *zap.SugaredLogger
	level  zapcore.Level
}

func Init(config *config.Logger, options ...zap.Option) (err error) {
	if logger != nil {
		return errors.New("logger already initialized")
	}

	logger, err = newLogger(config, options...)
	return
}

func New(config *config.Logger, options ...zap.Option) (*Logger, error) {
	return newLogger(config, options...)
}

func newLogger(config *config.Logger, options ...zap.Option) (*Logger, error) {
	hook := lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	w := zapcore.AddSync(&hook)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	l := level(config.Level)
	zapLevel := l.ZapLevel()
	var core zapcore.Core
	if zapLevel == zapcore.DebugLevel {
		core = zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.Lock(os.Stdout),
				zapLevel,
			),
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				w,
				zapLevel,
			),
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			w,
			zapLevel,
		)
	}

	zapLog := zap.New(core, options...)
	sugar := zapLog.Sugar()
	return &Logger{
		writer: zapLog,
		sugar:  sugar,
		config: config,
		level:  zapLevel,
	}, nil
}
func (l *Logger) Debug(str string, args ...zap.Field) {
	l.writer.Debug(str, args...)
}

func (l *Logger) Info(str string, args ...zap.Field) {
	l.writer.Info(str, args...)
}

func (l *Logger) Warn(str string, args ...zap.Field) {
	l.writer.Warn(str, args...)
}

func (l *Logger) Error(str string, args ...zap.Field) {
	l.writer.Error(str, args...)
}

func (l *Logger) Debugf(str string, args ...interface{}) {
	if l.level > zapcore.DebugLevel {
		return
	}
	l.sugar.Debugf(fmt.Sprintf(str, args...))
}
func (l *Logger) Infof(str string, args ...interface{}) {
	if l.level > zapcore.InfoLevel {
		return
	}
	l.sugar.Infof(fmt.Sprintf(str, args...))
}
func (l *Logger) Warnf(str string, args ...interface{}) {
	if l.level > zapcore.WarnLevel {
		return
	}
	l.sugar.Warnf(fmt.Sprintf(str, args...))
}
func (l *Logger) Errorf(str string, args ...interface{}) {
	if l.level > zapcore.ErrorLevel {
		return
	}
	l.sugar.Errorf(str, args...)
}

func (l *Logger) Sync() error {
	return l.writer.Sync()
}
