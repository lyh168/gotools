package logger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	Filename   string
	Level      string
	IsRotation bool
	Rotation   Rotation
}

type Rotation struct {
	MaxAge time.Duration
	Time   time.Duration
	Size   int64
	Count  uint
}

var _Config Config

func Configuration() Config {
	return _Config
}

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)
