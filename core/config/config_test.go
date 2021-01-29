package config

import (
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetReportCaller(true)
	os.Exit(m.Run())
}

func TestGetConfigPath(t *testing.T) {
	configPath := GetConfigPath()
	log.Info(configPath)
	SetConfigPath("/a/b/c.yml")
	configPath = GetConfigPath()
	log.Info(configPath)
	GetConfig()
}

