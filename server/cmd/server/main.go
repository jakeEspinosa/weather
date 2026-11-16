package main

import (
	"log/slog"
	"net/http"
	"os"
	"server/internal/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	api.NewServerMetricsHandler(logger).RegisterRoutes(mux)

	logger.Info("listening", "addr", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("server", "err", err)
		os.Exit(1)
	}
}
