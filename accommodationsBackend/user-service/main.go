package main

import (
	"accommodationsBackend/user-service/startup"
	"accommodationsBackend/user-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
