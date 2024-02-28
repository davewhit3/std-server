package server

import (
	"log/slog"
	"net/http"
)

func logRequest(handler http.Handler, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(
			"request incoming",
			slog.String("remoteAddr", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()))
		handler.ServeHTTP(w, r)
	})
}
