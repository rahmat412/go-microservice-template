package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/rahmat412/go-microservice-template/internal/app/server"
	"github.com/rahmat412/go-microservice-template/internal/util"
	"github.com/rahmat412/go-toolbox/container"
)

func main() {
	// Initialize application
	ctx, cancel, cfg, logger, err := util.InitializeApp()
	if err != nil {
		panic(fmt.Sprintf("Initialization error: %s", err.Error()))
	}
	defer cancel()

	logger.Info("Starting go-microservice-template service...")

	// Run migration
	err = server.RunMigration(cfg, logger)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error running migration: %s", err.Error()))
	}

	var wg sync.WaitGroup
	// Start the Chi server
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Info("[API] Starting API server...")
		apiServer := server.NewChiServer(cfg, logger)
		appContainer := container.New(apiServer, container.WithHTTPServer([]*http.Server{apiServer.HTTPServer()}))
		err := appContainer.Start(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("Server error: %s", err.Error()))
			cancel()
		} else {
			logger.Info("[API] API server is running.")
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
	logger.Info("[API] Application shutdown complete")
}
