package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"` // trace, debug, info, warn, error, fatal, panic, disabled
	HTTPPort int    `yaml:"http_port" env:"HTTP_PORT"`
	GRPCPort int    `yaml:"grpc_port" env:"GRPC_PORT"`
}

func LoadConfig(configFilePath string) (Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(configFilePath, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("error load config: %w", err)
	}
	return cfg, nil
}
