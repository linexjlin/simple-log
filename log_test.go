package log

import (
	"testing"
	"time"
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

func TestLogToWebsocket(t *testing.T) {
	go func() {
		for {
			log.Info("1")
			time.Sleep(time.Second * 2)
		}
	}()
	LogToWs(":8034", "/ws")
	select {}
}
