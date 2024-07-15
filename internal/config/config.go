package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config Конфиги grpc сервера, бд и прочего. Можно задавать через env, можно в yml конфиге
type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	GRPC
	Postgres
}

type GRPC struct {
	Host string `yaml:"host" env:"GRPC_HOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"GRPC_PORT" env-required:"true"`
}

type Postgres struct {
	Host     string `yaml:"host" env:"PG_HOST" env-default:"0.0.0.0"`
	Port     int    `yaml:"port" env:"PG_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"PG_USER" env-required:"true"`
	Password string `yaml:"password" env:"PG_PWD" env-required:"true"`
	Dbname   string `yaml:"dbname" env:"PG_DBNAME" env-default:"users"`
}

func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

const LocalConfigPath = "./local.yaml"

func MustLoad() Config {

	var cfg Config

	errEnv := cleanenv.ReadEnv(&cfg)
	//если из окружения не получили нужные параметры, пробуем взять конфиг файл
	if errEnv != nil {
		cfgPath := os.Getenv("CONFIG_PATH")

		if cfgPath == "" {
			if _, err := os.Stat(LocalConfigPath); os.IsNotExist(err) {
				log.Fatalf("config path not set and env reading error: %v", errEnv)
			}

			cfgPath = LocalConfigPath
		}

		if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
			log.Fatalf("config file not exists: %s", cfgPath)
		}

		err := cleanenv.ReadConfig(cfgPath, &cfg)
		if err != nil {
			log.Fatalf("failed to read config file: %s", err)
		}
	}

	return cfg
}
