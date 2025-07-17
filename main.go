package main

import (
	"log"
	"log/slog"

	"github.com/meesooqa/go-web-example/ext/demo"
	"github.com/meesooqa/go-web-example/ext/index"

	"github.com/meesooqa/go-web-example/cfg"
	"github.com/meesooqa/go-web-example/lgr"
	"github.com/meesooqa/go-web-example/srv"
	"github.com/meesooqa/go-web-example/srv/handlers"
	"github.com/meesooqa/go-web-example/srv/middlewares"
	"github.com/meesooqa/go-web-example/srv/theme"
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

	hh := []srv.Handler{
		handlers.NewStatic(logger, thm),
		index.New(logger, thm),
		demo.New(logger, thm),
	}
	mw := []srv.Middleware{
		middlewares.NewLogging(logger),
	}
	s := srv.New(conf.Server, hh, mw)

	logger.Info("server started", slog.String("host", conf.Server.Host()), slog.Int("port", conf.Server.Port()))
	log.Fatal(s.Run())
}
