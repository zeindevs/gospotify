package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url, err := h.Auth.Login(h.cfg.CLIENT_ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) HandleClientLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res, err := h.Auth.ClientLogin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"data": res})
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var code = r.URL.Query().Get("code")
	var state = r.URL.Query().Get("state")

	if state == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": "state required"})
		return
	}

	res, err := h.Auth.Callback(code, state)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
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
		log.Println("with secret.json")
		file, err := os.ReadFile("secret.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
			return
		}
		json.Unmarshal(file, &secret)
	} else {
		log.Println("with cookies")
		secret.RefreshToken = token.Value
	}

	res, err := h.Auth.RefreshToken(secret.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
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
