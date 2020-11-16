package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func UnmarshalLevel(level string) (*zap.AtomicLevel, error) {
	if level == "" {
		level = "debug"
	}
	zapLevel := zap.NewAtomicLevel()
	err := zapLevel.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}
	return &zapLevel, nil
}

func NewZapLogger(config Config, level *zap.AtomicLevel) (*zap.Logger, error) {
	encoder := zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			FunctionKey:    "func",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		})
	ws, err := getWriteSyncer(config)
	if err != nil {
		return nil, err
	}
	return zap.New(zapcore.NewTee(zapcore.NewCore(encoder, ws, level)), zap.AddCaller()), nil
}

func getWriteSyncer(config Config) (zapcore.WriteSyncer, error) {
	if len(config.Filename) == 0 {
		return os.Stdout, nil
	}
	if !config.IsRotation {
		file, _, err := zap.Open(config.Filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	dir := filepath.Dir(config.Filename)
	ext := filepath.Ext(config.Filename)
	filename := strings.TrimRight(filepath.Base(config.Filename), ext)
	timePattern := "%Y%m%d%H%M%S"
	writer, err := rotatelogs.New(
		filepath.Join(dir, fmt.Sprintf("%s.%s%s", filename, timePattern, ext)),
		rotatelogs.WithMaxAge(config.Rotation.MaxAge),
		rotatelogs.WithRotationTime(config.Rotation.Time),
		rotatelogs.WithRotationSize(config.Rotation.Size),
		rotatelogs.WithRotationCount(config.Rotation.Count),
		rotatelogs.WithLinkName(filepath.Join(dir, fmt.Sprintf("%s.%s%s", filename, timePattern, ext))),
		rotatelogs.ForceNewFile(),
	)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(writer), nil
}
