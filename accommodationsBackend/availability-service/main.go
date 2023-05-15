package main

import (
	"accommodationsBackend/availability-service/startup"
	"accommodationsBackend/availability-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
