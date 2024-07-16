package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const localConfigPath = "./local.yaml"

// Config Конфиги grpc сервера, бд и прочего. Можно задавать через env, можно в yml конфиге
type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	GRPC
	Postgres
}

// GRPC настройки grpc сервера
type GRPC struct {
	Host string `yaml:"host" env:"GRPC_HOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"GRPC_PORT" env-required:"true"`
}

// Postgres настройки подключения в бд
type Postgres struct {
	Host     string `yaml:"host" env:"PG_HOST" env-default:"0.0.0.0"`
	Port     int    `yaml:"port" env:"PG_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"PG_USER" env-required:"true"`
	Password string `yaml:"password" env:"PG_PWD" env-required:"true"`
	Dbname   string `yaml:"dbname" env:"PG_DBNAME" env-default:"users"`
}

// DSN генерирует строку подключения
func (p Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

// MustLoad загружает конфиг из окружения/файла. Фаталится если не получится
func MustLoad() Config {

	var cfg Config

	errEnv := cleanenv.ReadEnv(&cfg)
	if errEnv == nil {
		return cfg
	}

	//если из окружения не получили нужные параметры, пробуем взять конфиг файл
	cfgPath := os.Getenv("CONFIG_PATH")

	if cfgPath == "" {
		if _, err := os.Stat(localConfigPath); os.IsNotExist(err) {
			log.Fatalf("config path not set and env reading error: %v", errEnv)
		}

		cfgPath = localConfigPath
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file not exists: %s", cfgPath)
	}

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	return cfg
}
