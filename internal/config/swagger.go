package config

import "net"

// Swagger параметры swagger сервера
type Swagger struct {
	Host string `yaml:"swagger_host" env:"SWAGGER_HOST" env-default:"0.0.0.0"`
	Port string `yaml:"swagger_port" env:"SWAGGER_PORT" env-default:"8081"`
}

// Address возвращает строку подключения
func (s Swagger) Address() string {
	return net.JoinHostPort(s.Host, s.Port)
}
