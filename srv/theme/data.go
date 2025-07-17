package theme

import "html/template"

type TemplateData struct {
	Content DataContent
	Page    *DataPage
	Site    *DataSite
}

type DataContent any

type DataPage struct {
	Lang        string
	Title       string
	Description string
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
	Attr     template.HTMLAttr
	Children []DataMenuItem
}

func (t *Theme) SiteData() *DataSite {
	return &DataSite{
		Title:     "Lisa",
		SubTitle:  "The Leaseholder",
		BuildYear: "2025",
		Menus:     t.menus(),
	}
}

func (t *Theme) menus() map[string]DataMenuItem {
	return map[string]DataMenuItem{
		"Main": DataMenuItem{
			Children: []DataMenuItem{
				DataMenuItem{
					Name: "Home",
					Href: "/",
					Attr: "title=\"title\"",
				},
				DataMenuItem{
					Name: "Demo",
					Href: "/demo",
				},
			},
		},
	}
}
