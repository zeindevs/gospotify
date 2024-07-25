package internal

import (
	"encoding/json"

	"github.com/zeindevs/gospotify/config"
)

type PlayerService struct {
	http *Http
	cfg  *config.Config
}

func NewPlayerService(cfg *config.Config) *PlayerService {
	return &PlayerService{
		http: NewHttp(),
		cfg:  cfg,
	}
}

func (ps *PlayerService) GetCurrentPlaying(secret string) (any, error) {
	ps.http.Header.Add("Authorization", "Bearer "+secret)

	res, err := ps.http.Get("https://api.spotify.com/v1/me/player/currently-playing?market=ID")
	if err != nil {
		return nil, err
	}

	var val any
	json.Unmarshal(res, &val)

	return val, nil
}

// TODO: https://api.spotify.com/v1/me/player/next
func (ps *PlayerService) Next(secret string) (any, error) {
	return nil, nil
}

// TODO: https://api.spotify.com/v1/me/player/previous
func (ps *PlayerService) Prev(secret string) (any, error) {
	return nil, nil
}
