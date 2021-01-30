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
//	SetConfigPath("/a/b/c.yml")
//	configPath = GetConfigPath()
	log.Info(configPath)
	GetConfig()
}

func TestGetConfig(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Fail()
	}
	log.Warnf("config :\n%v\n\n", config)
	if config.DigitalVaults[0].Name != "test 1" {
		t.Errorf("Test fails. %v", config.DigitalVaults[0].Name) 
	}	
}
