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
	Styles    []template.HTML
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
		Styles:    t.styles(),
	}
}

func (t *Theme) menus() map[string]DataMenuItem {
	RegisterMenu(map[string]DataMenuItem{
		MainMenu: {
			Children: []DataMenuItem{{
				Sort: 100,
				Name: "Home",
				Href: "/",
				Attr: "title=\"title\"",
			}},
		},
	})
	menu := mergeMenu(menuRegistry...)
	for key, item := range menu {
		// TODO sortMenu is not working
		sortMenu(item.Children)
		menu[key] = item
	}
	return menu
}

func (t *Theme) styles() []template.HTML {
	RegisterCSS("<link rel=\"stylesheet\" href=\"/static/styles/styles.css\">")
	return cssRegistry
}
