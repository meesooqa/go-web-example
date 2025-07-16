package theme

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Config interface {
	ThemesDir() string
	Theme() string
}

type Theme struct {
	cfg Config
}

func New(cfg Config) *Theme {
	return &Theme{cfg: cfg}
}

func (t *Theme) BuildTemplate(contentFile string) (*template.Template, error) {
	if contentFile == "" {
		contentFile = "default.html"
	}
	layout := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", "main.html")
	page := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "content", contentFile)
	return template.ParseFiles(layout, page)
}

func (t *Theme) HandleStatic(mux *http.ServeMux) {
	path := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "assets")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}
