package demo

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/meesooqa/go-web-example/srv/handlers"
	"github.com/meesooqa/go-web-example/srv/theme"
)

type Demo struct {
	logger *slog.Logger
	t      handlers.Theme
}

func New(logger *slog.Logger, t handlers.Theme) *Demo {
	return &Demo{
		logger: logger,
		t:      t,
	}
}

func (h *Demo) Handle(mux *http.ServeMux) {
	mux.HandleFunc("/demo", h.handlePage)
}

func (h *Demo) handlePage(w http.ResponseWriter, r *http.Request) {
	//if r.Method != h.Method {
	//	http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	tmpl := h.t.MustBuildTemplate("demo.html", "")
	err := tmpl.Execute(w, h.data(r))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Demo) data(_ *http.Request) any {
	content := struct {
		DemoVar template.HTML
	}{
		DemoVar: template.HTML("<pre>Demo Var Value</pre>"),
	}
	data := &theme.TemplateData{
		Content: &content,
		Page: &theme.DataPage{
			Lang:        "en",
			Title:       "Demo",
			Description: "This is a demo page",
		},
		Site: h.t.SiteData(),
	}
	return data
}
