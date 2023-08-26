package main

import (
	"accommodationsBackend/event-store/startup"
	"accommodationsBackend/event-store/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
