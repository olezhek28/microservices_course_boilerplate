package config

import "net"

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
