package logger

import (
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
)

func NewLogger(cfg *config.Config) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(cfg.Log.LogLevel)
	if err != nil {
		return nil, errors.New("invalid log level")
	}

	zerolog.SetGlobalLevel(level)

	return &log.Logger, nil
}
