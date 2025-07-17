package theme

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Config interface {
	ThemesDir() string
	Theme() string
	// Dir() string
	// Name() string
	// Lang() string
}

type Theme struct {
	cfg Config
}

func New(cfg Config) *Theme {
	return &Theme{cfg: cfg}
}

func (t *Theme) MustBuildTemplate(content, layout string) *template.Template {
	if content == "" {
		content = "default.html"
	}
	if layout == "" {
		layout = "main.html"
	}

	layoutPath := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", layout)
	contentPath := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "content", content)
	tmpl := template.Must(template.ParseFiles(layoutPath, contentPath))

	layoutPartialsPattern := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", "inc", "*.html")
	tmpl = template.Must(tmpl.ParseGlob(layoutPartialsPattern))

	return tmpl
}

func (t *Theme) HandleStatic(mux *http.ServeMux) {
	path := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "assets")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}
