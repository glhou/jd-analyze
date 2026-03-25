package middleware

import (
	"log/slog"
	"net/http"
)

func loggerMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving", "url", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
