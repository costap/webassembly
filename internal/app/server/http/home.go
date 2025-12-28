package http

import (
	"github.com/costap/webassembly/internal/app/server/http/templates"
	"html/template"
	"net/http"
)

var index *template.Template

func init() {

	index = template.Must(template.New("index").Parse(string(templates.IndexTemplate)))
}

type HomeHandler struct {
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/index.html", "/":
		w.Header().Set("Content-Type", "text/html")
		if err := index.Execute(w, make(map[string]string)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
