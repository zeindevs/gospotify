package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, types.ApiResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		})
		return
	}

	data, err := h.PlayerService.GetCurrentPlaying(auth.AccessToken, h.cfg.MARKET)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, types.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := h.PlayerService.IsSaved(auth.AccessToken, data.Item.ID)
	for _, sv := range res {
		if sv {
			data.IsSaved = sv
			break
		}
	}

	WriteJSON(w, http.StatusOK, types.ApiResponse{
		Status: http.StatusOK,
		Data:   data,
	})
}

func (h *Handler) HandlePlayPrev(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, types.ApiResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		})
		return
	}

	res, err := h.PlayerService.Prev(auth.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, types.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	WriteJSON(w, http.StatusOK, types.ApiResponse{
		Status: http.StatusOK,
		Data:   res,
	})
}

func (h *Handler) HandlePlayNext(w http.ResponseWriter, r *http.Request) {
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, types.ApiResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		})
		return
	}

	res, err := h.PlayerService.Next(auth.AccessToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, types.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	WriteJSON(w, http.StatusOK, types.ApiResponse{
		Status: http.StatusOK,
		Data:   res,
	})
}

func (h *Handler) HandleSave(w http.ResponseWriter, r *http.Request) {
	var err error
	auth, err := GetAuth(r)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, types.ApiResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		})
		return
	}

	var req types.SaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusInternalServerError, types.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	var res any
	if !req.IsSaved {
		res, err = h.PlayerService.Save(auth.AccessToken, req)
	} else {
		res, err = h.PlayerService.DeleteSaved(auth.AccessToken, req)
	}
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, types.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	WriteJSON(w, http.StatusOK, types.ApiResponse{
		Status: http.StatusOK,
		Data:   res,
	})
}
