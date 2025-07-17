package index

import "github.com/meesooqa/go-web-example/srv/theme"

var mainMenu = map[string]theme.DataMenuItem{
	theme.MainMenu: {
		Children: []theme.DataMenuItem{
			{
				Sort: 100,
				Name: "Home",
				Href: "/",
				Attr: "title=\"Home\"",
			},
		},
	},
}
