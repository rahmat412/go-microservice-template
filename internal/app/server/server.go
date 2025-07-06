package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rahmat412/go-microservice-template/internal/config"
	"github.com/rahmat412/go-microservice-template/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/rahmat412/go-toolbox/logging"
)

type AppServer struct {
	srv    *http.Server
	router chi.Router
	cfg    *config.Config
	logger *logging.Logger

	internalConnection InternalConnection
}

func (appSrv *AppServer) BeforeStart(ctx context.Context) error {
	return appSrv.Run(ctx)
}

func (appSrv *AppServer) AfterStart(ctx context.Context) error {
	return appSrv.Stop(ctx)
}

func NewChiServer(cfg *config.Config, log *logging.Logger) *AppServer {
	router := chi.NewMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppHTTPPort),
		Handler: router,
	}

	return &AppServer{
		cfg:    cfg,
		router: router,
		logger: log,
		srv:    srv,
	}
}

func (appSrv *AppServer) Run(ctx context.Context) error {
	internalConnection := NewInternalConnection(appSrv.cfg, appSrv.logger)
	repository := NewRepository(internalConnection)
	service := NewService(repository, appSrv.cfg, appSrv.logger)
	newValidator := validator.New()

	userHandler := handler.NewUserHandler(appSrv.logger, service.UserService, newValidator)
	userHandler.RegisterRoutes(appSrv.router)

	appSrv.internalConnection = internalConnection

	return nil
}

func (appSrv *AppServer) Stop(ctx context.Context) error {
	if err := appSrv.internalConnection.Close(); err != nil {
		return err
	}

	return nil
}

func (appSrv *AppServer) HTTPServer() *http.Server {
	return appSrv.srv
}
