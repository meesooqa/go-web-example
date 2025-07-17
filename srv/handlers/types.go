package handlers

import (
	"html/template"
	"net/http"

	"github.com/meesooqa/go-web-example/srv/theme"
)

type StaticHandler interface {
	HandleStatic(mux *http.ServeMux)
	HandleStaticExt(ext string, mux *http.ServeMux)
}

type TemplateBuilder interface {
	MustBuildTemplateExt(ext, content, layout string) *template.Template
}

type DataSiteProvider interface {
	SiteData() *theme.DataSite
}

type Theme interface {
	StaticHandler
	TemplateBuilder
	DataSiteProvider
}
