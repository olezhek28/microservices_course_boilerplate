package config

import "net"

// GRPC настройки grpc сервера
type GRPC struct {
	Host string `yaml:"host" env:"GRPC_HOST" env-default:"0.0.0.0"`
	Port string `yaml:"port" env:"GRPC_PORT" env-required:"true"`
}

// Address адрес подключения
func (g GRPC) Address() string {
	return net.JoinHostPort(g.Host, g.Port)
}
