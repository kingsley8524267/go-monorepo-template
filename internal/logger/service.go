package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Debug(str string, args ...zap.Field) {
	logger.writer.Debug(str, args...)
}

func Info(str string, args ...zap.Field) {
	logger.writer.Info(str, args...)
}

func Warn(str string, args ...zap.Field) {
	logger.writer.Warn(str, args...)
}

func Error(str string, args ...zap.Field) {
	logger.writer.Error(str, args...)
}

func Debugf(str string, args ...interface{}) {
	if logger.level > zapcore.DebugLevel {
		return
	}
	logger.sugar.Debugf(fmt.Sprintf(str, args...))
}
func Infof(str string, args ...interface{}) {
	if logger.level > zapcore.InfoLevel {
		return
	}
	logger.sugar.Infof(fmt.Sprintf(str, args...))
}
func Warnf(str string, args ...interface{}) {
	if logger.level > zapcore.WarnLevel {
		return
	}
	logger.sugar.Warnf(fmt.Sprintf(str, args...))
}
func Errorf(str string, args ...interface{}) {
	if logger.level > zapcore.ErrorLevel {
		return
	}
	logger.sugar.Errorf(str, args...)
}

func Sync() error {
	return logger.writer.Sync()
}
