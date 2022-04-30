package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level uint8

const (
	LevelEmergency Level = iota
	LevelError
	LevelWarning
	LevelInformational
	LevelDebug
)

// Legacy log level constants to ensure backwards compatibility.
const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

// NewLogLevel is a convenience function that creates an AtomicLevel
func NewLogLevel(level Level) zap.AtomicLevel {
	var l zapcore.Level
	switch level {
	case LevelError:
		l = zapcore.ErrorLevel
	case LevelWarning:
		l = zapcore.WarnLevel
	case LevelInformational:
		l = zapcore.InfoLevel
	default: // default debug mode
		l = zapcore.DebugLevel
	}

	return zap.NewAtomicLevelAt(l)
}
