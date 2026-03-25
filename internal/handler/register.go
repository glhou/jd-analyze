package handler

import (
	"log/slog"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, logger *slog.Logger) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /{$}", index(logger))
}
