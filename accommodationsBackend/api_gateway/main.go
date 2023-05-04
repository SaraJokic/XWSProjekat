package main

import (
	"accommodationsBackend/api_gateway/startup"
	"accommodationsBackend/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
