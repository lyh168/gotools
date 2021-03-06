package log

var l *Logger

func init() {
	InitLogger(Config{})
}

func Default() *Logger {
	return l
}

func InitLogger(config Config) error {
	_Config = config
	var err error
	l, err = New(config)
	if err != nil {
		return err
	}
	return nil
}

func WithTraceID(traceID string) *Logger {
	logger := l.Clone()
	l.SetTraceID(traceID)
	return logger
}
