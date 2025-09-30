package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/flovan30/chi-boilerplate-api/internal/config"
	"github.com/flovan30/chi-boilerplate-api/internal/entity"
)

type Database struct {
	Gorm  *gorm.DB
	sqlDB *sql.DB
}

func InitDatabase(cfg *config.Config, log *zerolog.Logger) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Paris",
		cfg.Database.DBHost,
		cfg.Database.DBUser,
		cfg.Database.DBPassword,
		cfg.Database.DBName,
		cfg.Database.DBPort,
	)

	var gormLoggerLevel logger.LogLevel

	if config.IsDev(cfg) {
		gormLoggerLevel = logger.Info
	} else {
		gormLoggerLevel = logger.Silent
	}

	gormLogger := logger.Default.LogMode(gormLoggerLevel)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         gormLogger,
		PrepareStmt:    true,
		TranslateError: true,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sql.DB from GORM")
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(cfg.Database.DBMaxOpenConns)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	log.Info().Msg("Connected to the database")

	log.Info().Msg("Running AutoMigrate...")
	if err := db.AutoMigrate(
		// Models
		&entity.Book{},
	); err != nil {
		log.Error().Err(err).Msg("AutoMigrate failed")
		return nil, err
	}
	log.Info().Msg("Migrations done")

	return &Database{
		Gorm:  db,
		sqlDB: sqlDB,
	}, nil
}

func (d *Database) Close() error {
	if d.sqlDB != nil {
		return d.sqlDB.Close()
	}
	return nil
}
