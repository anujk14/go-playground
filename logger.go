package main

import (
	"log/slog"
	"os"
	"time"
)

func getLogger() *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.Int64Value(time.Now().Unix())
			}

			return a
		},
	})

	logger := slog.New(logHandler)
	return logger
}
