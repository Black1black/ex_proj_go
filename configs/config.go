package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// type DatabaseConfig struct {
//     Host     string `yaml:"host"`
//     Port     int `yaml:"port"`
//     User     string `yaml:"user"`
//     Password string `yaml:"password"`
//     DBName   string `yaml:"dbname"`
// }

// type Config struct {
//     Database DatabaseConfig `yaml:"database"`
// }

type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"postgres"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return &Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return &Config{}, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}
