package main

import (
	"accommodationsBackend/reservation-service/startup"
	"accommodationsBackend/reservation-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
