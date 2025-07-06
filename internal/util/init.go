package util

import (
	"context"
	"fmt"
	"os"

	"github.com/rahmat412/go-microservice-template/internal/config"
	"github.com/rahmat412/go-toolbox/logging"
)

func InitializeApp() (context.Context, context.CancelFunc, *config.Config, *logging.Logger, error) {
	// Handle shutdown signals
	ctx, cancel := HandleShutdownSignal(context.Background())

	// Load configuration
	cfg, err := config.GetConfig()
	if err != nil {
		cancel()
		return nil, nil, nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize logger
	logh := logging.NewHandler(os.Stderr,
		logging.WithAddSource(func(lvl logging.Level) bool {
			return lvl == cfg.GetLogLevel()
		}),
	)
	logger := logging.New(logh)

	return ctx, cancel, cfg, logger, nil
}
