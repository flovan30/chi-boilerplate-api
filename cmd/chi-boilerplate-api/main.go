package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/flovan30/chi-boilerplate-api/internal/app"
	"github.com/flovan30/chi-boilerplate-api/internal/config"
	"github.com/flovan30/chi-boilerplate-api/internal/database"
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

	db, err := database.InitDatabase(cfg, customLogger)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := app.NewRouter(customLogger, cfg, db.Gorm)

	startServerWithGracefulShutdown(router, cfg, db, customLogger)
}

func startServerWithGracefulShutdown(handler chi.Router, cfg *config.Config, db *database.Database, logger *zerolog.Logger) {
	addr := fmt.Sprintf("%s:%d", cfg.Api.AppHost, cfg.Api.AppPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		logger.Info().Msg(fmt.Sprintf("Server listening on %s", addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Msg("Listen error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Forced shutdown")
	}

	if err := db.Close(); err != nil {
		logger.Error().Err(err).Msg("Failed to close database")
	}

	logger.Info().Msg("Server and database shutdown cleanly")
}
