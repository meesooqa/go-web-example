package handlers

import (
	"log/slog"
	"net/http"
)

type Demo struct {
	logger *slog.Logger
	tb     TemplateBuilder
}

func NewDemo(logger *slog.Logger, tb TemplateBuilder) *Demo {
	return &Demo{
		logger: logger,
		tb:     tb,
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
	tmpl, err := h.tb.BuildTemplate("demo.html", "")
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, h.data(r))
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Demo) data(_ *http.Request) any {
	data := struct {
		Site    *DataSite
		Page    *DataPage
		Title   string
		DemoVar string
	}{
		Site: &DataSite{
			Title:     "Lisa",
			SubTitle:  "The Leaseholder",
			BuildYear: "2025",
			Menus: map[string]DataMenuItem{
				"Main": DataMenuItem{
					Children: []DataMenuItem{
						DataMenuItem{
							Name: "Home",
							Href: "/",
						},
						DataMenuItem{
							Name: "Demo",
							Href: "/demo",
						},
					},
				},
			},
		},
		Page: &DataPage{
			Lang:        "en",
			Title:       "Demo",
			Description: "This is a demo page",
		},
		DemoVar: "DemoVar <pre>VALUE</pre>",
	}
	return data
}
