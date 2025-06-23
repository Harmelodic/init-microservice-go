package main

import (
	"log/slog"
	"os"
)

func getLogger() *slog.Logger {
	loggerOutput := os.Getenv("LOGGER_OUTPUT")

	var handler slog.Handler

	switch loggerOutput {
	case "JSON":
		handler = slog.NewJSONHandler(os.Stdout, nil)
	case "CONSOLE":
		handler = slog.NewTextHandler(os.Stdout, nil)
	default:
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}

	return slog.New(handler)
}
