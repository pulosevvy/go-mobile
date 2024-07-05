package main

import (
	"go-mobile/config"
	"go-mobile/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
