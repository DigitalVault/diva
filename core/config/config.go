package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/DigitalVault/diva/core/info"
	"gopkg.in/yaml.v3"
)

var configPath string

type DigitalVaultType struct {
	Name string `yaml:"name"`
	DisplayName string `yaml:"displayname"`
	Path string `yaml:"path"`
}

type Config struct {
	DigitalVaults []DigitalVaultType `yaml:"digitalvaults"`
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

func WriteConfig(config Config, overwrite bool) error {
	log.Warnf("Opening %s", configPath)

	d, err := yaml.Marshal(&config)

	log.Warn(string(d))

	dir, _ := filepath.Split(configPath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(configPath, d, 0644)

	return err
}

func GetDefaultConfig() Config {
	data := `
digitalvaults:
 - name: "test 1"
   displayname: "disp 1"
   path: "path 1"
 - name: "test 2"
   displayname: "disp 2"
   path: "path 2"
`
	config := &Config{}
	err := yaml.Unmarshal([]byte(data), config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Infof("config:\n%v\n\n", config)
	return *config
}
