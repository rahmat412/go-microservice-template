package server

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rahmat412/go-microservice-template/internal/config"
	"github.com/rahmat412/go-toolbox/logging"
)

type InternalConnection struct {
	Db *sql.DB
}

func NewInternalConnection(cfg *config.Config, log *logging.Logger) InternalConnection {
	db, err := newDatabaseConnection(cfg.Database.ConnURL())
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	return InternalConnection{
		Db: db,
	}
}

func (i InternalConnection) Close() error {
	if err := i.Db.Close(); err != nil {
		return err
	}

	return nil
}

func newDatabaseConnection(connUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
