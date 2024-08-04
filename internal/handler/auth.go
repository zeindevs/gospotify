package handler

import (
	"net/http"
	"time"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url, err := h.AuthService.Login(h.cfg.CLIENT_ID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": err.Error()})
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) HandleClientLogin(w http.ResponseWriter, r *http.Request) {
	res, err := h.AuthService.ClientLogin()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"data": res})
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	var code = r.URL.Query().Get("code")
	var state = r.URL.Query().Get("state")

	if state == "" {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": "state required"})
		return
	}

	res, err := h.AuthService.Callback(code, state)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "AccessToken",
		Value:   res.AccessToken,
		Expires: time.Now().Add(60 * time.Minute),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "RefreshToken",
		Value:   res.RefreshToken,
		Expires: time.Now().Add(60 * time.Minute),
	})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	var secret types.AuthResponse
	token, err := r.Cookie("RefreshToken")
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": err.Error()})
		return
	} else {
		secret.RefreshToken = token.Value
	}

	res, err := h.AuthService.RefreshToken(secret.RefreshToken)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]any{"err": err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "AccessToken",
		Value:   res.AccessToken,
		Expires: time.Now().Add(60 * time.Minute),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "RefreshToken",
		Value:   res.RefreshToken,
		Expires: time.Now().Add(60 * time.Minute),
	})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "AccessToken",
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "RefreshToken",
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
