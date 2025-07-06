package config

import (
	"fmt"
	"strings"

	"github.com/rahmat412/go-toolbox/config"
	"github.com/rahmat412/go-toolbox/logging"
)

type DatabaseConfig struct {
	Host            string `env:"DB_HOST"`
	Port            string `env:"DB_PORT"`
	User            string `env:"DB_USER"`
	Password        string `env:"DB_PASSWORD"`
	Name            string `env:"DB_NAME"`
	EnableMigration bool   `env:"DB_ENABLE_MIGRATION"`
	SSLMode         string `env:"DB_SSL_MODE"`
}

type Config struct {
	LogLevel    string `env:"LOG_LEVEL"`
	AppHTTPPort string `env:"APP_HTTP_PORT" default:"8080"`
	Port        int    `env:"PORT"`
	Database    DatabaseConfig
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	err := config.LoadEnv(&cfg)
	if err != nil {
		return nil, nil
	}

	return &cfg, nil
}

func (c *Config) GetLogLevel() logging.Level {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return logging.LevelDebug
	case "info":
		return logging.LevelInfo
	case "warn":
		return logging.LevelWarn
	case "error":
		return logging.LevelError
	case "fatal":
		return logging.LevelFatal
	default:
		return logging.LevelInfo
	}
}

func (db DatabaseConfig) ConnURL() string {
	f := "postgres://%s:%s@%s:%s/%s?sslmode=%s"
	connUrl := fmt.Sprintf(f, db.User, db.Password, db.Host, db.Port, db.Name, db.SSLMode)

	return connUrl
}
