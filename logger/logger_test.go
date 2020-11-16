package logger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := New(Config{
		Filename:   "/tmp/gotools.log",
		Level:      "debug",
		IsRotation: true,
		Rotation: Rotation{
			MaxAge: time.Minute,
			Time:   time.Second * 10,
			Size:   1024 * 1024 * 8,
		},
	})
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
