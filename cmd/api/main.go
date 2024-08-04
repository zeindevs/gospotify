package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zeindevs/gospotify/internal/config"
	"github.com/zeindevs/gospotify/internal/handler"
	"github.com/zeindevs/gospotify/internal/middleware"
	"github.com/zeindevs/gospotify/internal/service"
	"github.com/zeindevs/gospotify/internal/util"
)

func main() {
	cfg := config.NewConfig()
	handler := handler.NewHandler(&handler.HandlerConfig{
		Config:        cfg,
		AuthService:   service.NewAuthService(cfg),
		PlayerService: service.NewPlayerService(cfg),
	})

	s := &http.ServeMux{}
	httpFS := http.FS(os.DirFS("public"))
	fileServer := http.FileServer(httpFS)

	s.Handle("GET /", middleware.LoggerAsset(fileServer))
	s.Handle("GET /login", middleware.Logger(handler.HandleLogin))
	s.Handle("GET /refresh", middleware.Logger(handler.HandleRefresh))
	s.Handle("GET /login/client", middleware.Logger(handler.HandleClientLogin))
	s.Handle("GET /logout", middleware.Logger(handler.HandleLogout))
	s.Handle("GET /callback", middleware.Logger(handler.HandleCallback))
	s.Handle("GET /api/playing", middleware.Logger(handler.HandlePlaying))
	s.Handle("POST /api/playing/next", middleware.Logger(handler.HandlePlayNext))
	s.Handle("POST /api/playing/prev", middleware.Logger(handler.HandlePlayPrev))
	s.Handle("POST /api/track/save", middleware.Logger(handler.HandleSave))

	log.Println("Server up and listening on http://localhost:9001")
	util.ErrorPanic(http.ListenAndServe(":9001", s))
}
