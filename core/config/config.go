package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/DigitalVault/diva/core/info"
	"gopkg.in/yaml.v3"
)

var configPath string

type ScannerType struct {
	Patterns []string `yaml:"patterns",flow`
}
type Config struct {
	Scanner ScannerType `yaml:"scanner"`
}

func GetConfigPath() string {
	if (configPath == "") {
		newConfigPath := fmt.Sprintf("%s/.diva/diva.yml", info.Info.HomeDir)
		log.Debugf("No config path defined. Defaults to %s", newConfigPath)
		configPath = newConfigPath
	}
	log.Debugf("Config path : %s", configPath)
  return configPath
}

func SetConfigPath(newConfigPath string) {
	configPath = newConfigPath
}

func GetConfig() (*Config, error) {
	
	configPath = GetConfigPath()
	config := &Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Infof("There is no config file at %s. Creating one with default settings.", configPath)
		// create default Config struct
		config := GetDefaultConfig()
		WriteConfig(config, true)
	}
	file, err := os.Open(configPath)
	if (err != nil) {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func WriteConfig(config Config, overwrite bool) {
}

func GetDefaultConfig() Config {
	config := &Config{
		Scanner : ScannerType {
			Patterns: []string{"*.jpg", "*.png"},
		},
	}
	log.Info(config)
	return *config
}
