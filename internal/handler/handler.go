package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/service"
)

type Handler struct {
	cfg           *config.Config
	AuthService   *service.AuthService
	PlayerService *service.PlayerService
}

type HandlerConfig struct {
	Config        *config.Config
	AuthService   *service.AuthService
	PlayerService *service.PlayerService
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		cfg:           cfg.Config,
		AuthService:   cfg.AuthService,
		PlayerService: cfg.PlayerService,
	}
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
