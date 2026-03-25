package handler

import (
	"log/slog"
	"net/http"

	"github.com/glhou/jd-analyze/views"
)

func index(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Getting index page")
		views.Index().Render(r.Context(), w)
	}
}
