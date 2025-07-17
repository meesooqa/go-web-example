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

func (t *Theme) BuildTemplate(content, layout string) (*template.Template, error) {
	if content == "" {
		content = "default.html"
	}
	if layout == "" {
		layout = "main.html"
	}

	layoutPath := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", layout)
	contentPath := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "content", content)

	return template.ParseFiles(layoutPath, contentPath)
}

func (t *Theme) HandleStatic(mux *http.ServeMux) {
	path := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "assets")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}
