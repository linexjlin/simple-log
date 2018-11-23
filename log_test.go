package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	log.Debugf("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Critical("Critical")
}

func TestLogToFile(t *testing.T) {
	DebugEanble(true)
	LogToFile("/tmp/example.txt")
	log.Debugf("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Critical("Critical")
}
