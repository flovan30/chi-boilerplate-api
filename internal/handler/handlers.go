package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
)

func RegisterRoutes(router chi.Router, logger *zerolog.Logger, cfg *config.Config, db *gorm.DB) {
	// init repositories

	// init services

	// init handlers

	router.Route("/api", func(r chi.Router) {

		r.Get("/alive", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("ok")
		})
	})
}
