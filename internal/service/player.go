package service

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/pkg"
	"github.com/zeindevs/gospotify/types"
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

func (ps *PlayerService) GetCurrentPlaying(secret, marketID string) (*types.CurrentPlayingResponse, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	res, _, err := ps.http.Get(fmt.Sprintf("https://api.spotify.com/v1/me/player/currently-playing?market=%s", marketID))
	if err != nil {
		return nil, err
	}

	var val types.CurrentPlayingResponse
	if err := json.Unmarshal(res, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

func (ps *PlayerService) Prev(secret string) (any, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	res, _, err := ps.http.Post("https://api.spotify.com/v1/me/player/previous", nil)
	if err != nil {
		return nil, err
	}

	var val any
	if err := json.Unmarshal(res, &val); err != nil {
		return nil, err
	}

	return val, nil
}

func (ps *PlayerService) Next(secret string) (any, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	res, _, err := ps.http.Post("https://api.spotify.com/v1/me/player/next", nil)
	if err != nil {
		return nil, err
	}

	var val any
	if err := json.Unmarshal(res, &val); err != nil {
		return nil, err
	}

	return val, nil
}

func (ps *PlayerService) Save(secret string, ids types.SaveRequest) (any, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	ps.http.Header.Set("Content-Type", "application/json")
	data, err := json.Marshal(map[string][]string{"ids": ids.IDs})
	if err != nil {
		return nil, err
	}
	_, _, err = ps.http.Put("https://api.spotify.com/v1/me/tracks", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	return "saved", nil
}

func (ps *PlayerService) IsSaved(secret, ids string) ([]bool, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	res, _, err := ps.http.Get(fmt.Sprintf("https://api.spotify.com/v1/me/tracks/contains?ids=%s", ids))
	if err != nil {
		return nil, err
	}

	var val []bool
	if err := json.Unmarshal(res, &val); err != nil {
		return nil, err
	}

	return val, nil
}

// TODO: DELETE https://api.spotify.com/v1/me/tracks --data '{"ids": []}'
func (ps *PlayerService) DeleteSaved(secret string, ids types.SaveRequest) (any, error) {
	ps.http.Header.Set("Authorization", "Bearer "+secret)
	ps.http.Header.Set("Content-Type", "application/json")
	data, err := json.Marshal(map[string][]string{"ids": ids.IDs})
	if err != nil {
		return nil, err
	}
	_, _, err = ps.http.Delete("https://api.spotify.com/v1/me/tracks", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	return "unsaved", nil
}
