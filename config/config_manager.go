package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Settings Struct

// Struct - define the config
type Struct struct {
	Player1 string `yaml:"player1"`
	Player2 string `yaml:"player2"`
}

func Load() {
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &Settings)
	if err != nil {
		panic(err)
	}
}
