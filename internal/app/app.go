package app

import "log/slog"

type App struct {
	Logger *slog.Logger
}

func New(logger *slog.Logger) *App {
	return &App{
		Logger: logger,
	}
}
