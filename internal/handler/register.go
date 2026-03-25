package handler

import (
	"log/slog"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, logger *slog.Logger) {
	mux.HandleFunc("GET /", index(logger))
}
