package handlers

import (
	"log/slog"
	"net/http"

	"github.com/meesooqa/go-web-example/server/theme"
)

type Index struct {
	logger *slog.Logger
	t      Theme
}

func NewIndex(logger *slog.Logger, t Theme) *Index {
	return &Index{
		logger: logger,
		t:      t,
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
	tmpl, err := h.t.BuildTemplate("", "")
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
	return &theme.TemplateData{
		Page: &theme.DataPage{
			Lang:        "en",
			Title:       "Index",
			Description: "This is a Home page",
		},
		Site: h.t.SiteData(),
	}
}
