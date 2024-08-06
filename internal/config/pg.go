package config

import "fmt"

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
