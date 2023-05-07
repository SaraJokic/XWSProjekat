package config

import "os"

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		UserDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		UserDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
