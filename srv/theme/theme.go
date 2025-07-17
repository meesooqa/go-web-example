package theme

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Config interface {
	ThemesDir() string
	Theme() string
	ExtDir() string
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

func (t *Theme) MustBuildTemplateExt(ext, content, layout string) *template.Template {
	if content == "" {
		content = "default.html"
	}
	if layout == "" {
		layout = "main.html"
	}

	contentPath := ""
	if ext != "" {
		contentPath = filepath.Join(t.cfg.ExtDir(), ext, t.cfg.ThemesDir(), t.cfg.Theme(), "content", content)
	} else {
		contentPath = filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "content", content)
	}
	layoutPath := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", layout)
	tmpl := template.Must(template.ParseFiles(layoutPath, contentPath))

	layoutPartialsPattern := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "layouts", "inc", "*.html")
	tmpl = template.Must(tmpl.ParseGlob(layoutPartialsPattern))

	return tmpl
}

func (t *Theme) HandleStatic(mux *http.ServeMux) {
	path := filepath.Join(t.cfg.ThemesDir(), t.cfg.Theme(), "assets")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}

func (t *Theme) HandleStaticExt(ext string, mux *http.ServeMux) {
	if ext == "" {
		t.HandleStatic(mux)
		return
	}
	path := filepath.Join(t.cfg.ExtDir(), ext, t.cfg.ThemesDir(), t.cfg.Theme(), "assets")
	ptrn := fmt.Sprintf("/ext/%s/static/", ext)
	mux.Handle(ptrn, http.StripPrefix(ptrn, http.FileServer(http.Dir(path))))
}
