package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string
	StoragePath string
	*HTTPServer
}

type HTTPServer struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

func MustLoad(path string, name string, _type string) (*Config, error) {

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(_type)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	server := &HTTPServer{Address: viper.GetString("http_server.address"),
		Timeout:     viper.GetDuration("http_server.timeout"),
		IdleTimeout: viper.GetDuration("http_server.idle_timeout")}

	return &Config{Env: viper.GetString("env"), StoragePath: viper.GetString("storage_path"), HTTPServer: server}, nil
}
