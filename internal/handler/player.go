package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	data, err := h.PlayerService.GetCurrentPlaying(auth.AccessToken, "ID")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	res, err := h.PlayerService.IsSaved(auth.AccessToken, data.Item.ID)
	isSaved := false
	for _, sv := range res {
		if sv {
			isSaved = true
		}
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": data, "is_saved": isSaved})
}

func (h *Handler) HandlePlayPrev(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	res, err := h.PlayerService.Prev(auth.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandlePlayNext(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	res, err := h.PlayerService.Next(auth.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandleSave(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	var req types.SaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	// res, err := h.PlayerService.IsSaved(auth.AccessToken, req.IDs)
	res, err := h.PlayerService.Save(auth.AccessToken, req)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}
