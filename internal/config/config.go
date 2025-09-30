package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Api      ApiConfig
	Log      LogConfig
	Database DBConfig
}

type ApiConfig struct {
	AppEnv  string
	AppHost string
	AppPort int
}

type LogConfig struct {
	LogLevel string
}

type DBConfig struct {
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         int
	DBSSLMode      string
	DBMaxOpenConns int
}

func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.ReadInConfig()

	// key | zero or undefined is allowed ?
	requiredVars := map[string]bool{
		"APP_ENV":                false,
		"APP_HOST":               false,
		"APP_PORT":               false,
		"LOG_LEVEL":              true,
		"DB_HOST":                false,
		"DB_USER":                false,
		"DB_PASSWORD":            false,
		"DB_NAME":                false,
		"DB_PORT":                false,
		"DB_SSL_MODE":            false,
		"DB_MAX_OPEN_CONNECTION": false,
	}

	for key, allowZero := range requiredVars {
		if !viper.IsSet(key) {
			return nil, fmt.Errorf("missing required environment variable: %s", key)
		}

		if !allowZero {
			if viper.GetString(key) == "" {
				return nil, fmt.Errorf("empty or invalid value for: %s", key)
			}
			if viper.GetInt(key) == 0 && strings.HasSuffix(key, "_PORT") {
				return nil, fmt.Errorf("invalid zero value for: %s", key)
			}
		}
	}

	cfg := &Config{
		Api: ApiConfig{
			AppEnv:  viper.GetString("APP_ENV"),
			AppHost: viper.GetString("APP_HOST"),
			AppPort: viper.GetInt("APP_PORT"),
		},
		Log: LogConfig{
			LogLevel: viper.GetString("LOG_LEVEL"),
		},
		Database: DBConfig{
			DBHost:         viper.GetString("DB_HOST"),
			DBUser:         viper.GetString("DB_USER"),
			DBPassword:     viper.GetString("DB_PASSWORD"),
			DBName:         viper.GetString("DB_NAME"),
			DBPort:         viper.GetInt("DB_PORT"),
			DBSSLMode:      viper.GetString("DB_SSL_MODE"),
			DBMaxOpenConns: viper.GetInt("DB_MAX_OPEN_CONNECTION"),
		},
	}

	return cfg, nil
}

func IsDev(cfg *Config) bool {
	return cfg.Api.AppEnv == "dev"
}
