package ctrl

import (
	"github.com/davewhit3/std-server/server"
	"net/http"
)

func SetupHealthHandlers(srv *http.ServeMux) {
	srv.HandleFunc("GET /health", health)
	srv.HandleFunc("GET /ready", ready)
}

func health(w http.ResponseWriter, r *http.Request) {
	server.Response(w).Code(http.StatusOK).JsonBody(map[string]bool{"status": true})
}

func ready(w http.ResponseWriter, r *http.Request) {
	server.Response(w).Code(http.StatusOK).JsonBody(map[string]bool{"status": true})
}
