package main

import (
	"log"
	"log/slog"

	"github.com/meesooqa/go-web-example/cfg"
	"github.com/meesooqa/go-web-example/lgr"
	"github.com/meesooqa/go-web-example/server"
	"github.com/meesooqa/go-web-example/server/handlers"
	"github.com/meesooqa/go-web-example/server/middlewares"
	"github.com/meesooqa/go-web-example/server/theme"
)

func main() {
	conf, err := cfg.Load("etc/config.yml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}
	logger, closer := lgr.New(conf.Log)
	if closer != nil {
		defer func() { _ = closer.Close() }()
	}

	thm := theme.New(conf.Theme)

	hh := []server.Handler{
		handlers.NewStatic(logger, thm),
		handlers.NewIndex(logger, thm),
		handlers.NewDemo(logger, thm),
	}
	mw := []server.Middleware{
		middlewares.NewLogging(logger),
	}
	s := server.New(conf.Server, hh, mw)

	logger.Info("server started", slog.String("host", conf.Server.Host()), slog.Int("port", conf.Server.Port()))
	log.Fatal(s.Run())
}
