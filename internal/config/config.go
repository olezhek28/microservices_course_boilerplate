package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const localConfigPath = "./.env"

// Config Конфиги grpc сервера, бд и прочего. Можно задавать через env, можно в yml конфиге
type Config struct {
	Env           string `yaml:"env" env:"ENV" env-required:"true"`
	UsersCacheTTL int    `yaml:"users_cache_ttl" env:"USERS_CACHE_TTL" env-default:"60"`
	GRPC
	Postgres
	Redis
	Kafka
	HTTP
	Swagger
	NewUsersTopic string `yaml:"new_users_topic" env:"NEW_USERS_TOPIC" env-required:"true"`
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
