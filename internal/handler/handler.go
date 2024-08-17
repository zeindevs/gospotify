package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/service"
	"github.com/zeindevs/gospotify/types"
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

func GetAuth(r *http.Request) (*types.AuthResponse, error) {
	var auth types.AuthResponse
	accessToken, err := r.Cookie("AccessToken")
	if err != nil {
		return nil, err
	}
	auth.AccessToken = accessToken.Value
	refreshToken, err := r.Cookie("RefreshToken")
	if err != nil {
		return nil, err
	}
	auth.RefreshToken = refreshToken.Value
	return &auth, nil
}
