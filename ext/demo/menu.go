package demo

import "github.com/meesooqa/go-web-example/srv/theme"

var mainMenu = map[string]theme.DataMenuItem{
	theme.MainMenu: {
		Children: []theme.DataMenuItem{
			{
				Sort: 200,
				Name: "Demo",
				Href: "/demo",
				Children: []theme.DataMenuItem{
					{
						Sort: 20,
						Name: "Sub Item 2",
						Href: "/demo/page2",
					},
					{
						Sort: 10,
						Name: "Sub Item 1",
						Href: "/demo/page1",
					},
				},
			},
		},
	},
}
