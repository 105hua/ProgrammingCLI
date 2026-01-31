package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

const DEFAULT_PATH string = "config.json"

type Config struct {
	ApiKey string `json:"api_key"`
	Model  string `json:"model"`
}

func GetDefaultConfig() Config {
	return Config{
		ApiKey: "YOUR_API_KEY",
		Model:  "minimax/minimax-m2.1",
	}
}

func SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("Error mashaling JSON: %w", err)

	}
	return os.WriteFile(DEFAULT_PATH, data, 0644)
}

func LoadConfig() Config {
	data, err := os.ReadFile(DEFAULT_PATH)
	if err != nil {
		defaultConfig := GetDefaultConfig()
		SaveConfig(defaultConfig)
		return defaultConfig
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return GetDefaultConfig()
	}

	return config
}
