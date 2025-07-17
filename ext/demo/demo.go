package demo

import "github.com/meesooqa/go-web-example/srv/theme"

const code = "demo"

func init() {
	theme.RegisterMenu(mainMenu)
}
