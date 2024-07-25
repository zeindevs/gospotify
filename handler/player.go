package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var secret types.AuthResponse

	token, err := r.Cookie("AccessToken")
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
		secret.AccessToken = token.Value
	}

	data, err := h.Player.GetCurrentPlaying(secret.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"data": data})
}
