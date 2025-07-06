package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rahmat412/go-microservice-template/internal/config"
	"github.com/rahmat412/go-toolbox/logging"
)

func RunMigration(cfg *config.Config, log *logging.Logger) error {
	if !cfg.Database.EnableMigration {
		return nil
	}

	dsn := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	dsn = fmt.Sprintf(dsn, cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error(err.Error())
	}
	defer db.Close()

	log.Info("Running database migration...")
	migrationDir := "./database/migrations"
	if err := goose.Up(db, migrationDir); err != nil {
		return err
	}
	log.Info("Database migration completed successfully.")

	return nil
}
