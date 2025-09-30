package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
	"github.com/flovan30/chi-boilerplate-api/internal/handler"
)

func NewRouter(logger *zerolog.Logger, cfg *config.Config, db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(
		middleware.AllowContentType("application/json"),
		middleware.CleanPath,
		middleware.GetHead,
		middleware.Logger,
		middleware.NoCache,
		middleware.StripSlashes,
		// middleware.Throttle(100), // limit the number of request globaly for all the api
		middleware.Timeout(10*time.Second),
		httprate.LimitByRealIP(500, 1*time.Minute),
		middleware.Recoverer,
	)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * int(time.Hour),
	}))

	handler.RegisterRoutes(r, logger, cfg, db)

	return r
}
