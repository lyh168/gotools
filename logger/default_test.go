package logger

import "testing"

type Person struct {
	Name string
	ID   int
}

func TestInitLogger(t *testing.T) {
	Default().Debug("Info", "qe23", "sdasdsad")
	Default().Debug(&Person{
		Name: "jax",
		ID:   10,
	})
	Default().Debugf("Info %d", 1)
	Default().Info("Info")
	Default().Infof("Info %d", 1)
	Default().Warn("Info")
	Default().Warnf("Info %d", 1)
	Default().Error("Info")
	Default().Errorf("Info %d", 1)
	Default().Panic("Info")
	Default().Panicf("Info %d", 1)
	Default().Fatal("Info")
	Default().Fatalf("Info %d", 1)
}
