package model

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Version  string `json:"version"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func LoadTextVenture(configFileName string) ServiceConfig {

	log.Printf("reading %s...", configFileName)

	// Read and parse the json configuration file
	configPath, _ := filepath.Abs(configFileName)
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}
	var serviceConfig ServiceConfig
	err = yaml.Unmarshal(configData, &serviceConfig)
	if err != nil {
		log.Fatal("Error parsing configuration file:", err)
	}

	// Use the configuration
	log.Println("Name from config file:", serviceConfig.Name)

	return serviceConfig
}
