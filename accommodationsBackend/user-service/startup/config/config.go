package config

import "os"

type Config struct {
	Port                            string
	UserDBHost                      string
	UserDBPort                      string
	NatsHost                        string
	NatsPort                        string
	NatsUser                        string
	NatsPass                        string
	CancelReservationCommandSubject string
	CancelReservationReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                            os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:                      os.Getenv("USER_DB_HOST"),
		UserDBPort:                      os.Getenv("USER_DB_PORT"),
		NatsHost:                        os.Getenv("NATS_HOST"),
		NatsPort:                        os.Getenv("NATS_PORT"),
		NatsUser:                        os.Getenv("NATS_USER"),
		NatsPass:                        os.Getenv("NATS_PASS"),
		CancelReservationCommandSubject: os.Getenv("CANCEL_RESERVATION_COMMAND_SUBJECT"),
		CancelReservationReplySubject:   os.Getenv("CANCEL_RESERVATION_REPLY_SUBJECT"),
	}
}
