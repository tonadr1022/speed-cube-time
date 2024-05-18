package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

const (
	defaultServerPort         = 8080
	defaultJWTExpirationHours = 72
)

// config for the server
type Config struct {
	// Data source name to connect to db. Required
	DSN string `yaml:"dsn" env:"DSN"`
	// JWT signing key. Required
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY"`
	// Server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// JWT expireation in hours. Defaults to 72 hours
	JWTExpirationHours int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
}

func Load(file string) (*Config, error) {
	// default config
	config := Config{
		ServerPort:         defaultServerPort,
		JWTExpirationHours: defaultJWTExpirationHours,
	}

	// load and parse config file
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	// load from env here
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
