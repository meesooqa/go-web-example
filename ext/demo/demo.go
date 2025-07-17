package demo

import "github.com/meesooqa/go-web-example/srv/theme"

const code = "demo"

func init() {
	theme.RegisterMenu(mainMenu)
	theme.RegisterCSS("<link rel=\"stylesheet\" href=\"/ext/demo/static/styles/styles.css\">")
}
