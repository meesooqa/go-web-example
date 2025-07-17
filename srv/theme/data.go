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
	Sort     int
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
		"Main": {
			Children: []DataMenuItem{
				{
					Name: "Home",
					Href: "/",
					Attr: "title=\"title\"",
				},
				{
					Name: "Demo",
					Href: "/demo",
					Children: []DataMenuItem{
						{
							Name: "Sub Item 1",
							Href: "/demo/page1",
						},
						{
							Name: "Sub Item 2",
							Href: "/demo/page2",
						},
					},
				},
			},
		},
	}
}
