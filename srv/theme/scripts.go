package theme

import "html/template"

var scriptRegistry = make([]template.HTML, 0)

func RegisterScript(css template.HTML) {
	scriptRegistry = append(scriptRegistry, css)
}
