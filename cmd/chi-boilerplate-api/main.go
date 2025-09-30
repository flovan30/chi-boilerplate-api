package main

import (
	"log"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
	"github.com/flovan30/chi-boilerplate-api/internal/logger"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	customLogger, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatal("Failed to initialise logger", err)
	}
}
