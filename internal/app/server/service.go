package server

import (
	"github.com/rahmat412/go-microservice-template/internal/config"
	"github.com/rahmat412/go-microservice-template/internal/service"
	"github.com/rahmat412/go-toolbox/logging"
)

type Service struct {
	UserService service.UserServiceProvider
}

func NewService(repo Repository, cfg *config.Config, log *logging.Logger) Service {
	return Service{
		UserService: service.NewUserService(
			repo.UserPostgresRepository,
			log,
		),
	}
}
