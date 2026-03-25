package middleware

import (
	"log/slog"
	"net/http"
)

// Chain apply global middelwares
func ApplyGlobalMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return loggerMiddleware(next, logger)
}
