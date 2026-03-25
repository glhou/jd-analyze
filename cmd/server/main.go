package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/glhou/jd-analyze/internal/handler"
	"github.com/glhou/jd-analyze/internal/middleware"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)
	slog.SetDefault(logger)
	mux := http.NewServeMux()

	handler.RegisterHandlers(mux, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Server running on", "port", port)
	err := http.ListenAndServe(":"+port, middleware.ApplyGlobalMiddleware(mux, logger))
	if err != nil {
		logger.Error("Error", "error", err.Error())
	}
}
