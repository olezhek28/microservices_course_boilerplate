package config

import "net"

// HTTP параметры http сервера
type HTTP struct {
	Host string `yaml:"http_host" env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port string `yaml:"http_port" env:"HTTP_PORT" env-default:"8080"`
}

// Address возвращает строку подключения
func (h HTTP) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}
