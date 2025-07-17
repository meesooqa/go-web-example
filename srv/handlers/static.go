package handlers

import (
	"log/slog"
	"net/http"
)

type Static struct {
	logger *slog.Logger
	sh     StaticHandler
}

func NewStatic(logger *slog.Logger, sh StaticHandler) *Static {
	return &Static{
		logger: logger,
		sh:     sh,
	}
}

func (h *Static) Handle(mux *http.ServeMux) {
	h.sh.HandleStatic(mux)
}
