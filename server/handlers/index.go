package handlers

import (
	"log/slog"
	"net/http"
)

type Index struct {
	logger *slog.Logger
	tb     TemplateBuilder
}

func NewIndex(logger *slog.Logger, tb TemplateBuilder) *Index {
	return &Index{
		logger: logger,
		tb:     tb,
	}
}

func (h *Index) Handle(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handlePage)
}

func (h *Index) handlePage(w http.ResponseWriter, r *http.Request) {
	//if r.Method != h.Method {
	//	http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	tmpl, err := h.tb.BuildTemplate("", "")
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

func (h *Index) data(_ *http.Request) any {
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
			Title:       "Index",
			Description: "This is a Home page",
		},
	}
	return data
}
