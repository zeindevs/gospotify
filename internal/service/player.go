package service

import (
	"encoding/json"
	"fmt"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/pkg"
)

type PlayerService struct {
	http *pkg.Http
	cfg  *config.Config
}

func NewPlayerService(cfg *config.Config) *PlayerService {
	return &PlayerService{
		http: pkg.NewHttp(),
		cfg:  cfg,
	}
}

func (ps *PlayerService) GetCurrentPlaying(secret, marketID string) (any, error) {
	ps.http.Header.Add("Authorization", "Bearer "+secret)
	res, err := ps.http.Get(fmt.Sprintf("https://api.spotify.com/v1/me/player/currently-playing?market=%s", marketID))
	if err != nil {
		return nil, err
	}

	var val any
	json.Unmarshal(res, &val)

	return val, nil
}

func (ps *PlayerService) Prev(secret string) (any, error) {
	ps.http.Header.Add("Authorization", "Bearer "+secret)
	res, err := ps.http.Post("https://api.spotify.com/v1/me/player/previous", nil)
	if err != nil {
		return nil, err
	}

	var val any
	json.Unmarshal(res, &val)

	return val, nil
}

func (ps *PlayerService) Next(secret string) (any, error) {
	ps.http.Header.Add("Authorization", "Bearer "+secret)
	res, err := ps.http.Post("https://api.spotify.com/v1/me/player/next", nil)
	if err != nil {
		return nil, err
	}

	var val any
	json.Unmarshal(res, &val)

	return val, nil
}

// TODO: PUT https://api.spotify.com/v1/me/tracks --data {'ids': ?}
func (ps *PlayerService) Save(secret, ids string) (any, error) {
	return nil, nil
}

// TODO: GET https://api.spotify.com/v1/me/tracks/contains?ids=
func (ps *PlayerService) IsSaved(secret, ids string) (any, error) {
	return nil, nil
}
