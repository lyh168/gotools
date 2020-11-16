package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	lv      *zap.AtomicLevel
	l       *zap.Logger
	sl      *zap.SugaredLogger
	traceID string
}

func New(config Config) (*Logger, error) {
	level, err := UnmarshalLevel(config.Level)
	if err != nil {
		return nil, err
	}
	zapLogger, err := NewZapLogger(config, level)
	if err != nil {
		return nil, err
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(1))
	return &Logger{
		lv: level,
		l:  zapLogger,
		sl: zapLogger.Sugar(),
	}, nil
}

func (this *Logger) Clone() *Logger {
	return &Logger{
		lv: this.lv,
		l:  this.l,
		sl: this.l.Sugar(),
	}
}

func (this *Logger) ZapLogger() *zap.Logger {
	return this.l
}

func (this *Logger) ZapSugaredLogger() *zap.SugaredLogger {
	return this.sl
}

func (this *Logger) Debug(v ...interface{}) {
	this.sl.Debug(v...)
}

func (this *Logger) Debugf(format string, v ...interface{}) {
	this.sl.Debugf(format, v...)
}

func (this *Logger) Info(v ...interface{}) {
	this.sl.Info(v...)
}

func (this *Logger) Infof(format string, v ...interface{}) {
	this.sl.Infof(format, v...)
}

func (this *Logger) Warn(v ...interface{}) {
	this.sl.Warn(v...)
}

func (this *Logger) Warnf(format string, v ...interface{}) {
	this.sl.Warnf(format, v...)
}

func (this *Logger) Error(v ...interface{}) {
	this.sl.Error(v...)
}

func (this *Logger) Errorf(format string, v ...interface{}) {
	this.sl.Errorf(format, v...)
}

func (this *Logger) Panic(v ...interface{}) {
	this.sl.Panic(v...)
}

func (this *Logger) Panicf(format string, v ...interface{}) {
	this.sl.Panicf(format, v...)
}

func (this *Logger) Fatal(v ...interface{}) {
	this.sl.Fatal(v...)
}

func (this *Logger) Fatalf(format string, v ...interface{}) {
	this.sl.Fatalf(format, v...)
}

func (this *Logger) SetTraceID(str string) {
	if len(str) > 0 {
		this.traceID = str
		this.sl.With(zap.String("traceId", this.traceID))
	}
}

func (this *Logger) GetTraceID() string {
	return this.traceID
}
