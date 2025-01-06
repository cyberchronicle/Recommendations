// config.go
package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration structure
type ApplicationConfig struct {
	MLHostPort              string `yaml:"ml_host_port"`
	BackendHostPort         string `yaml:"backend_host_port"`
	PersonalAccountHostPort string `yaml:"personal_account_host_port"`
	RecsDbHostPort          string `yaml:"recs_db_host_port"`
	ScrapperHostPort        string `yaml:"scrapper_host_port"`
}

// Global variable to hold the configuration
var AppConfig ApplicationConfig

// LoadConfig loads configuration from a YAML file
func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	return nil
}
