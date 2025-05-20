package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Settings Struct

// Struct - define the config
type Struct struct {
	Player1 string `yaml:"player1"`
	Player2 string `yaml:"player2"`
}

func Load() error {
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return fmt.Errorf("error reading the config file: %w", err)
	}

	err = yaml.Unmarshal(configFile, &Settings)
	if err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
