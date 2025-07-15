package main

import (
	"log"

	"github.com/meesooqa/go-web-example/config"
	"github.com/meesooqa/go-web-example/logging"
	"github.com/meesooqa/go-web-example/server"
	"github.com/meesooqa/go-web-example/server/handlers"
	"github.com/meesooqa/go-web-example/server/middlewares"
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

	hh := []server.Handler{
		handlers.NewStatic(logger, cfg.Handler),
		handlers.NewIndex(logger, cfg.Handler),
	}
	mw := []server.Middleware{
		middlewares.NewLogging(logger),
	}
	s := server.New(cfg.Server, hh, mw)

	log.Printf("Сервер запущен на http://%s:%d", cfg.Server.Host(), cfg.Server.Port())
	log.Fatal(s.Run())
}
