package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func HandleShutdownSignal(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	return ctx, cancel
}
