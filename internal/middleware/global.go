package middleware

import (
	"log/slog"
	"net/http"
)

// Apply global middlewares in a chain
func ApplyGlobalMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return loggerMiddleware(next, logger)
}
