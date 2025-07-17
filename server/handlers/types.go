package handlers

import "html/template"

type TemplateBuilder interface {
	BuildTemplate(content, layout string) (*template.Template, error)
}

type DataSite struct {
	Title     string
	SubTitle  string
	BuildYear string
	Menus     map[string]DataMenuItem
}

type DataMenuItem struct {
	Name     string
	Href     string
	Attr     string
	Children []DataMenuItem
}

type DataPage struct {
	Lang        string
	Title       string
	Description string
}
