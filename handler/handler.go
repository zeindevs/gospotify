package handler

import (
	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/internal"
)

type Handler struct {
	cfg    *config.Config
	Auth   *internal.AuthService
	Player *internal.PlayerService
}

func NewHandler(cfg *config.Config, auth *internal.AuthService, player *internal.PlayerService) *Handler {
	return &Handler{
		cfg:    cfg,
		Auth:   auth,
		Player: player,
	}
}
