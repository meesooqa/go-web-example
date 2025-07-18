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
	Scripts   []template.HTML
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
		Scripts:   t.scripts(),
	}
}

func (t *Theme) menus() map[string]DataMenuItem {
	menu := mergeMenu(menuRegistry...)
	for key, item := range menu {
		sortMenu(item.Children)
		menu[key] = item
	}
	return menu
}

func (t *Theme) styles() []template.HTML {
	// theme common css
	RegisterCSS("<link rel=\"stylesheet\" href=\"/static/styles/styles.css\">")
	return cssRegistry
}

func (t *Theme) scripts() []template.HTML {
	// theme common js
	RegisterScript("<script type=\"module\" src=\"/static/scripts/index.js\"></script>")
	return scriptRegistry
}
