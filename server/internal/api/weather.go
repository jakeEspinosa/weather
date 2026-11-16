package api

import (
	"log/slog"
	"net/http"
	"server/internal/lib"
)

type ServerMetricsHandler struct {
	logger *slog.Logger
}

func NewServerMetricsHandler(l *slog.Logger) *ServerMetricsHandler {
	return &ServerMetricsHandler{logger: l}
}

func (h *ServerMetricsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/ping", h.pingHandler)
}

func (h *ServerMetricsHandler) pingHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("pinged")
	lib.WriteJson(w, http.StatusOK, `{"msg": "pong"}`)
}
