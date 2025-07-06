package server

import (
	"github.com/rahmat412/go-microservice-template/internal/repository/pgsql"
)

type Repository struct {
	UserPostgresRepository pgsql.UserRepositoryProvider
}

func NewRepository(client InternalConnection) Repository {
	return Repository{
		UserPostgresRepository: pgsql.NewUserRepository(client.Db),
	}
}
