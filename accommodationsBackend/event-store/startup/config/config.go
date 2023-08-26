package config

import "os"

type Config struct {
	Port         string
	EventsDBHost string
	EventsDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:         os.Getenv("EVENT_STORE_PORT"),
		EventsDBHost: os.Getenv("EVENTS_DB_HOST"),
		EventsDBPort: os.Getenv("EVENTS_DB_PORT"),
	}
}
