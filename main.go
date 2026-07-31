package main

import (
	"Go-Ecom-Aws/config"
	"Go-Ecom-Aws/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded %v\n", err)
	}

	api.StartServer(cfg)
}
