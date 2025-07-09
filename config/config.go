package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Username string `json:"username"`
}

var (
	configFileName = "config.json"
	configDirName  = "cfcli"
)

// SaveConfig saves config to ~/.config/cfcli/config.json
func SaveConfig(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadConfig reads config from file
func LoadConfig() (Config, error) {
	var cfg Config

	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func getConfigFilePath() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, configDirName, configFileName), nil
}
