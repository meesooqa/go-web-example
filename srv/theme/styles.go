package theme

import "html/template"

var cssRegistry = make([]template.HTML, 0)

func RegisterCSS(css template.HTML) {
	cssRegistry = append(cssRegistry, css)
}
