package util

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	Logger.Debug("this is debug log")
	Logger.Info("this is info log")
	Logger.Warn("this is warn log")
	Logger.Error("this is error log")
	Logger.Fatal("this is fatal log")
}
