package commons

import (
	"log/slog"
	"os"
)

// NewLogger makes a new sensibly preconfigured slog.Logger for use in an application.
// The log format will be JSON or plain text depending on the value of the `LOG_FORMAT` environment variable.
func NewLogger() *slog.Logger {
	loggerOutput := os.Getenv("LOG_FORMAT")

	var handler slog.Handler

	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}

	switch loggerOutput {
	case "JSON":
		handler = slog.NewJSONHandler(os.Stdout, &handlerOptions)
	case "TEXT":
		handler = slog.NewTextHandler(os.Stdout, &handlerOptions)
	default:
		handler = slog.NewJSONHandler(os.Stdout, &handlerOptions)
	}

	return slog.New(handler)
}
