package main

import (
	"accommodationsBackend/auth-service/startup"
	"accommodationsBackend/auth-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
