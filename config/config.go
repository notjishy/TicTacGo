package config

import (
	"fmt"
	"github.com/TwiN/go-color"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
)

// default config settings for when config does not yet exist
const (
	defaults = `# player settings
# supported colors: black, red, green, yellow, blue, purple, cyan, grey, gray, white
player1:
  symbol: X
  color: red
player2:
    symbol: O
    color: cyan`
)

var path string // path to config file

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
	// get config path based on user's OS
	switch runtime.GOOS {
	case "windows":
		// %APPDATA%
		path = filepath.Join(os.Getenv("APPDATA"), "TicTacGo")
	case "darwin": // macOS
		// ~/Library/Application Support/
		homeDir, _ := os.UserHomeDir()
		path = filepath.Join(homeDir, "Library", "Application Support", "TicTacGo")
	default: // Linux
		// ~/.config/
		path = filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "TicTacGo")
		if path == filepath.Join("", "TicTacGo") {
			// XDG_CONFIG_HOME not set, fallback to ~/.config
			homeDir, _ := os.UserHomeDir()
			path = filepath.Join(homeDir, ".config", "TicTacGo")
		}
	}

	configPath := filepath.Join(path, "config.yaml")
	file, err := os.ReadFile(configPath)
	if err != nil {
		// create config file if it doesn't exist already
		if os.IsNotExist(err) {
			// make directory path
			if err := os.MkdirAll(path, 0755); err != nil {
				return fmt.Errorf("error creating config directory: %w", err)
			}

			// write config file
			if err := os.WriteFile(configPath, []byte(defaults), 0644); err != nil {
				return fmt.Errorf("error creating config file: %w", err)
			}

			// update with default config
			file = []byte(defaults)
		} else { // when config already exists
			return fmt.Errorf("error reading the config file: %w", err)
		}
	}

	if err := yaml.Unmarshal(file, &Settings); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
