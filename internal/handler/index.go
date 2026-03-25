package handler

import (
	"log/slog"
	"net/http"

	"github.com/glhou/jd-analyze/views"
)

func index(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := views.Index().Render(r.Context(), w); err != nil {
			logger.Error("failed to render index page",
				slog.Any("err", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
