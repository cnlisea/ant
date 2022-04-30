package logs

var DefaultLogger *Logger

func Fatal(s string, fields ...Field) {
	if DefaultLogger != nil {
		DefaultLogger.Fatal(s, fields...)
	}
}

func Err(s string, fields ...Field) {
	if DefaultLogger != nil {
		DefaultLogger.Error(s, fields...)
	}
}

func Warn(s string, fields ...Field) {
	if DefaultLogger != nil {
		DefaultLogger.Warn(s, fields...)
	}
}

func Info(s string, fields ...Field) {
	if DefaultLogger != nil {
		DefaultLogger.Info(s, fields...)
	}
}

func Debug(s string, fields ...Field) {
	if DefaultLogger != nil {
		DefaultLogger.Debug(s, fields...)
	}
}

func Close() {
	if DefaultLogger != nil {
		DefaultLogger.Close()
	}
}

func Init(path string, level Level, encodingJson bool, callerSkip ...int) error {
	var (
		err  error
		skip int // skip
	)

	// caller skip
	if callerSkip != nil {
		for i := range callerSkip {
			skip += callerSkip[i]
		}
	}

	encoding := EncodingConsole
	if encodingJson {
		encoding = EncodingJson
	}

	DefaultLogger, err = NewLogger(encoding, path, level, skip+1)

	return err
}
