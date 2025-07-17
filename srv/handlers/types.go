package handlers

import (
	"html/template"

	"github.com/meesooqa/go-web-example/srv/theme"
)

type TemplateBuilder interface {
	MustBuildTemplateExt(ext, content, layout string) *template.Template
}

type DataSiteProvider interface {
	SiteData() *theme.DataSite
}

type Theme interface {
	TemplateBuilder
	DataSiteProvider
}
