package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type level string

func (l *level) String() string {
	ls := strings.ToLower(string(*l))
	switch ls {
	case "debug", "info", "warn", "error":
		return ls
	default:
		return ""
	}
}

func (l *level) ZapLevel() zapcore.Level {
	ls := l.String()
	switch ls {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.DebugLevel
	}
}
