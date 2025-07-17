package index

import "github.com/meesooqa/go-web-example/srv/theme"

const code = "index"

func init() {
	theme.RegisterMenu(mainMenu)
	theme.RegisterCSS("<link rel=\"stylesheet\" href=\"/ext/index/static/styles/styles.css\">")
}
