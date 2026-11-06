package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// Запуск сервиса
func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	<-ctx.Done()
	return nil
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("service failed", "err", err)
		os.Exit(1)
	}
}
