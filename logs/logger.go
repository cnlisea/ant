package logs

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Log    *zap.Logger
	closed bool
	mutex  sync.Mutex
}

func NewLogger(encoding Encoding, path string, level Level, callerSkip int) (*Logger, error) {
	cfg := &zap.Config{
		Level:            NewLogLevel(level),
		Encoding:         NewEncoding(encoding),
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{path},
		ErrorOutputPaths: []string{path},
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err := cfg.Build(zap.AddCallerSkip(1 + callerSkip))

	return &Logger{Log: log}, err
}

func (l *Logger) Close() {
	l.mutex.Lock()
	if !l.closed {
		l.closed = true
		if l.Log != nil {
			l.Log.Sync()
		}
	}
	l.mutex.Unlock()
}

func (l *Logger) CloneWithCallerSkip(callerSkip int) *Logger {
	return &Logger{
		Log: l.Log.WithOptions(zap.AddCallerSkip(callerSkip)),
	}
}

func (l *Logger) Fatal(s string, fields ...Field) {
	l.Log.Fatal(s, fields...)
}

func (l *Logger) Error(s string, fields ...Field) {
	l.Log.Error(s, fields...)
}

func (l *Logger) Warn(s string, fields ...Field) {
	l.Log.Warn(s, fields...)
}

func (l *Logger) Info(s string, fields ...Field) {
	l.Log.Info(s, fields...)
}

func (l *Logger) Debug(s string, fields ...Field) {
	l.Log.Debug(s, fields...)
}
