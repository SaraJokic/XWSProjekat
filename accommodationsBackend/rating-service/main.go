package main

import (
	"accommodationsBackend/rating-service/startup"
	"accommodationsBackend/rating-service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
