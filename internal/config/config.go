package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	v *viper.Viper
}

type ServerConfig struct {
	Server     string
	GrpcServer string
}

func NewConfig(v *viper.Viper) *Configuration {
	return &Configuration{
		v: v,
	}
}

func (c *Configuration) GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Server:     c.v.GetString("service.port"),
		GrpcServer: c.v.GetString("service.grpc-port"),
	}
}

func (c *Configuration) GetDatabseConfig(name string) *DatabaseConfig {
	return &DatabaseConfig{
		host:     c.v.GetString(fmt.Sprintf("database.%s.host", name)),
		port:     c.v.GetInt(fmt.Sprintf("database.%s.port", name)),
		user:     c.v.GetString(fmt.Sprintf("database.%s.usert", name)),
		password: c.v.GetString(fmt.Sprintf("database.%s.password", name)),
		schema:   c.v.GetString(fmt.Sprintf("database.%s.schema", name)),
		DatabasePoolConfig: DatabasePoolConfig{
			maxCon:            c.v.GetInt(fmt.Sprintf("database.%s.max-conn", name)),
			minCon:            c.v.GetInt(fmt.Sprintf("database.%s.min-conn", name)),
			keepAliveInterval: c.v.GetDuration(fmt.Sprintf("database.%s.keep-alive-interval", name)),
			maxConnsLifeTime:  c.v.GetDuration(fmt.Sprintf("database.%s.max-conn-lifetime", name)),
		},
	}
}
