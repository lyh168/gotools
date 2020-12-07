package log

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNewZapLogger(t *testing.T) {
	lv := zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := NewZapLogger(Config{
		Filename:   "/tmp/gotools.log",
		Level:      "debug",
		IsRotation: true,
		Rotation: Rotation{
			MaxAge: time.Minute,
			Time:   time.Second * 10,
			Size:   1024 * 1024 * 8,
		},
	}, &lv)
	if err != nil {
		t.Fatal(err)
	}
	c := time.After(2 * time.Minute)
OUT:
	for {
		select {
		case <-c:
			break OUT
		default:
			logger.Info("hello")
		}
	}

}
