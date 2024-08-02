package config

import (
	"fmt"
	"log"
	"net"
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

// Redis настройки подключения в бд
type Redis struct {
	Host              string `yaml:"host" env:"REDIS_HOST" env-default:"0.0.0.0"`
	Port              string `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
	ConnectionTimeout int    `yaml:"connection_timeout" env:"REDIS_CONTIME" env-default:"5"`
	IdleTimeout       int    `yaml:"idle_timeout" env:"REDIS_IDLETIME" env-default:"300"`
	MaxIdle           int    `yaml:"max_idle" env:"REDIS_MAXIDLE" env-default:"10"`
}

// Address адрес подключения
func (r Redis) Address() string {
	return net.JoinHostPort(r.Host, r.Port)
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
