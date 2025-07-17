package handlers

import (
	"html/template"

	"github.com/meesooqa/go-web-example/srv/theme"
)

type TemplateBuilder interface {
	BuildTemplate(content, layout string) (*template.Template, error)
}

type DataSiteProvider interface {
	SiteData() *theme.DataSite
}

type Theme interface {
	TemplateBuilder
	DataSiteProvider
}
