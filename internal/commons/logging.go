package commons

import (
	"io"
	"log/slog"
)

type LogFormat string

const LogFormatJSON LogFormat = "JSON"
const LogFormatTEXT LogFormat = "TEXT"

// NewLogger makes a new sensibly preconfigured slog.Logger for use in an application.
// The log format will be LogFormatJSON or plain text depending on the value of the `LOG_FORMAT` environment variable.
func NewLogger(format LogFormat, writer io.Writer) *slog.Logger {
	var handler slog.Handler

	handlerOptions := slog.HandlerOptions{
		AddSource: true,
	}

	switch format {
	case LogFormatJSON:
		handler = slog.NewJSONHandler(writer, &handlerOptions)
	case LogFormatTEXT:
		handler = slog.NewTextHandler(writer, &handlerOptions)
	default:
		handler = slog.NewJSONHandler(writer, &handlerOptions)
	}

	return slog.New(handler)
}
