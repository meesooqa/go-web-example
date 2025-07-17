package main

import (
	"log"
	"log/slog"

	"github.com/meesooqa/go-web-example/config"
	"github.com/meesooqa/go-web-example/logging"
	"github.com/meesooqa/go-web-example/server"
	"github.com/meesooqa/go-web-example/server/handlers"
	"github.com/meesooqa/go-web-example/server/middlewares"
	"github.com/meesooqa/go-web-example/server/theme"
)

func main() {
	cfg, err := config.Load("etc/config.yml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}
	logger, closer := logging.New(cfg.Log)
	if closer != nil {
		defer func() { _ = closer.Close() }()
	}

	thm := theme.New(cfg.Theme)

	hh := []server.Handler{
		handlers.NewStatic(logger, thm),
		handlers.NewIndex(logger, thm),
		handlers.NewDemo(logger, thm),
	}
	mw := []server.Middleware{
		middlewares.NewLogging(logger),
	}
	s := server.New(cfg.Server, hh, mw)

	logger.Info("server started", slog.String("host", cfg.Server.Host()), slog.Int("port", cfg.Server.Port()))
	log.Fatal(s.Run())
}
