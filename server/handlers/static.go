package handlers

import (
	"log/slog"
	"net/http"
	"path/filepath"
)

type Static struct {
	logger *slog.Logger
	cfg    Config
}

func NewStatic(logger *slog.Logger, cfg Config) *Static {
	return &Static{
		logger: logger,
		cfg:    cfg,
	}
}

func (h *Static) Handle(mux *http.ServeMux) {
	path := filepath.Join(h.cfg.TemplatesDir(), h.cfg.TemplateName(), "static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}
