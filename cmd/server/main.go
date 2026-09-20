package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quachnamluong/website-ke-toa/internal/app"
	"github.com/quachnamluong/website-ke-toa/internal/server"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	a := app.New(logger)
	s := server.New(a, ":8080")

	go func() {
		if err := s.Start(); err != nil {
			logger.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", "err", err)
		os.Exit(1)
	}

	logger.Info("server stopped cleanly")
}
