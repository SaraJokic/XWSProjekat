package main

import (
	"accommodationsBackend/accommodations-service/startup"
	"accommodationsBackend/accommodations-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
