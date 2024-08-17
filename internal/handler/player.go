package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	secret, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	data, err := h.PlayerService.GetCurrentPlaying(secret.AccessToken, "ID")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": data})
}

func (h *Handler) HandlePlayPrev(w http.ResponseWriter, r *http.Request) {
	secret, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	res, err := h.PlayerService.Prev(secret.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandlePlayNext(w http.ResponseWriter, r *http.Request) {
	secret, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	res, err := h.PlayerService.Next(secret.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandleSave(w http.ResponseWriter, r *http.Request) {
	_, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
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
