package http

import (
	"github.com/costap/webassembly/internal/app/server/http/resources"
	"net/http"
)

func init() {

}

type AssetsHandler struct {
}

func (h *AssetsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/assets/js/wasm_exec.js":
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(resources.WasmExecJs)
	case "/assets/wasm/app.wasm":
		w.Header().Set("Content-Type", "application/wasm")
		w.Write(resources.AppWasm)
	default:
		http.NotFound(w, r)
	}
}
