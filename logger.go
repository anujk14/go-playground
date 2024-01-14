package main

import (
	"log/slog"
	"os"
)

func getLogger() *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(logHandler)
	return logger
}
