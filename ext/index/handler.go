package index

import (
	"log/slog"
	"net/http"

	"github.com/meesooqa/go-web-example/srv/handlers"
	"github.com/meesooqa/go-web-example/srv/theme"
)

type Index struct {
	logger *slog.Logger
	t      handlers.Theme

	route string
}

func New(logger *slog.Logger, t handlers.Theme) *Index {
	return &Index{
		logger: logger,
		t:      t,
		route:  "/",
	}
}

func (h *Index) Handle(mux *http.ServeMux) {
	h.handleStatic(mux)
	mux.HandleFunc(h.route, h.handlePage)
}

func (h *Index) handleStatic(mux *http.ServeMux) {
	h.t.HandleStaticExt(code, mux)
}

func (h *Index) handlePage(w http.ResponseWriter, r *http.Request) {
	//if r.Method != h.Method {
	//	http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	tmpl := h.t.MustBuildTemplateExt(code, "index.html", "")
	err := tmpl.Execute(w, h.data(r))
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
