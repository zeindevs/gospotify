package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	var secret types.AuthResponse
	token, err := r.Cookie("AccessToken")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	} else {
		secret.AccessToken = token.Value
	}

	data, err := h.PlayerService.GetCurrentPlaying(secret.AccessToken, "ID")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": data})
}

func (h *Handler) HandlePlayPrev(w http.ResponseWriter, r *http.Request) {
	var secret types.AuthResponse
	token, err := r.Cookie("AccessToken")
	if err != nil {

		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	} else {
		secret.AccessToken = token.Value
	}

	res, err := h.PlayerService.Prev(secret.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandlePlayNext(w http.ResponseWriter, r *http.Request) {
	var secret types.AuthResponse
	token, err := r.Cookie("AccessToken")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	} else {
		secret.AccessToken = token.Value
	}

	res, err := h.PlayerService.Next(secret.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandleSave(w http.ResponseWriter, r *http.Request) {
	var secret types.AuthResponse
	token, err := r.Cookie("AccessToken")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	} else {
		secret.AccessToken = token.Value
	}

	type likeRequest struct {
		IDs string `json:"ids"`
	}
	var req likeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusInternalServerError, map[string]any{"data": nil})
}
