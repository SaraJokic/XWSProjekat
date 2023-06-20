package config

import "os"

type Config struct {
	Port          string
	RatingsDBHost string
	RatingsDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("RATING_SERVICE_PORT"),
		RatingsDBHost: os.Getenv("RATING_DB_HOST"),
		RatingsDBPort: os.Getenv("RATING_DB_PORT"),
	}
}
