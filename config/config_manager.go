package config

import (
	"fmt"
	"github.com/TwiN/go-color"
	"gopkg.in/yaml.v3"
	"os"
)

var Settings Struct

// Player - player specific settings
type Player struct {
	Symbol string `yaml:"symbol"`
	Color  string `yaml:"color"`
}

// Struct - define the config structure
type Struct struct {
	Player1 Player `yaml:"player1"`
	Player2 Player `yaml:"player2"`
}

// GetColor - return translated color from color name in config
func (p Player) GetColor() string {
	// all supported colors (minus white as it is default)
	colorMap := map[string]string{
		"red":    color.Red,
		"cyan":   color.Cyan,
		"black":  color.Black,
		"green":  color.Green,
		"yellow": color.Yellow,
		"blue":   color.Blue,
		"purple": color.Purple,
		"gray":   color.Gray,
		"grey":   color.Gray,
	}

	// verify color is valid and return it
	if c, exists := colorMap[p.Color]; exists {
		return c
	}

	return color.White // default
}

// Load - load the config file
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
