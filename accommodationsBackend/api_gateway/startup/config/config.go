package config

import "os"

type Config struct {
	Port              string
	UserHost          string
	UserPort          string
	AccommodationHost string
	AccommodationPort string
	AuthHost          string
	AuthPort          string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AuthHost:          os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:          os.Getenv("AUTH_SERVICE_PORT"),
	}
}
