package logging

import (
	"io"
	"log"
	"log/slog"
	"os"
)

//go:generate moq --out log_mock.go . Config

type Config interface {
	IsWriteToFile() bool
	Path() string
	Level() slog.Level
	OutputFormat() string
}

func New(cfg Config) (*slog.Logger, io.Closer) {
	if cfg.IsWriteToFile() {
		return fileLogger(cfg)
	}
	return consoleLogger(cfg), nil
}

func consoleLogger(cfg Config) *slog.Logger {
	return slog.New(handler(cfg, os.Stdout))
}

// logger, closer := fileLogger(cfg)
// defer closer.Close()
func fileLogger(cfg Config) (*slog.Logger, io.Closer) {
	file, err := os.OpenFile(cfg.Path(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	logger := slog.New(handler(cfg, file))
	return logger, file
}

func handler(cfg Config, w io.Writer) slog.Handler {
	opts := &slog.HandlerOptions{
		Level: cfg.Level(),
	}
	switch cfg.OutputFormat() {
	case "text":
		return slog.NewTextHandler(w, opts)
	default:
		return slog.NewJSONHandler(w, opts)
	}
}
