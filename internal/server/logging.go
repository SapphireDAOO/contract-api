package server

import (
	"log/slog"
	"os"
	"strings"
)

// setupLogging installs the process-wide logger. Text reads better in a
// terminal, so production emits JSON for a log shipper and everything else
// stays human-readable. LOG_LEVEL selects verbosity and defaults to info.
func setupLogging() {
	opts := &slog.HandlerOptions{Level: logLevel()}

	var handler slog.Handler
	if _, production := os.LookupEnv("PRODUCTION"); production {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}

func logLevel() slog.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
