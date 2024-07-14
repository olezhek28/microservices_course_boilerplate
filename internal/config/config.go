package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// GRPCConfig Конфиги grpc сервера, можно задавать через env, можно в yml конфиге
type GRPCConfig struct {
	Host string `yaml:"host" env:"GHOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"GPORT" env-default:"50500"`
}

func MustLoad() GRPCConfig {
	var cfg GRPCConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("failed to load config env: %v", err)
	}

	return cfg
}
