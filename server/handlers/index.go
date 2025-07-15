package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"
)

type Index struct {
	logger *slog.Logger
	cfg    Config
}

func NewIndex(logger *slog.Logger, cfg Config) *Index {
	return &Index{
		logger: logger,
		cfg:    cfg,
	}
}

func (h *Index) Handle(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handlePage)
}

func (h *Index) handlePage(w http.ResponseWriter, r *http.Request) {
	// TODO "pages/index.html"
	fn := filepath.Join(h.cfg.TemplatesDir(), h.cfg.TemplateName(), "layout.html")
	tmpl, err := template.ParseFiles(fn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Hello, World!",
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
