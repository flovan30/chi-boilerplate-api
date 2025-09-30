package main

import (
	"log"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
}
