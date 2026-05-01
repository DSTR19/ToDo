package core_http_server

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Addr            string        `envconfig:"SERVER_ADDR" required:"true"`
	ShutDownTimeout time.Duration `envconfig:"SERVER_SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("Get HTTP Server Config: %w", err)
		panic(err)
	}
	return config
}
