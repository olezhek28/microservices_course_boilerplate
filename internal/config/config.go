package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig struct {
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"PORT" env-default:"50500"`
}

func MustLoad() GRPCConfig {
	var cfg GRPCConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("failed to load config env: %v", err)
	}

	return cfg
}
