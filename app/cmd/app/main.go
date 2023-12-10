package main

import (
	"app/internal/config"
	"app/pkg/logging"
	"log"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := logging.GetLogger()

	app, err := app.newApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
}
