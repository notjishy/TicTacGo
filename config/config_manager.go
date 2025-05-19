package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var config Struct

// Struct - define the config
type Struct struct {
	Player1 string `yaml:"player1"`
	Player2 string `yaml:"player2"`
}

func GetConfig() Struct {
	// only read file if it hasnt already been read and loaded
	if config.Player1 == "" {
		loadConfig()
	}
	return config
}

func loadConfig() {
	fmt.Println("read config")
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
}
