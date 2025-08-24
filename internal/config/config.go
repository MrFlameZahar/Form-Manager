package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

func NewConfigFromFile(configPath string) (*Config, error) {
	if configPath == "" {
		return &Config{}, fmt.Errorf("config path is empty")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("config file is not exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &Config{}, fmt.Errorf("cant read config")
	}
	return &cfg, nil
}
