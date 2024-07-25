package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/handler"
	"github.com/zeindevs/gospotify/internal"
)

func intercept404(handler, on404 http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hookedWriter := &hookedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(hookedWriter, r)

		if hookedWriter.got404 {
			on404.ServeHTTP(w, r)
		}
	})
}

type hookedResponseWriter struct {
	http.ResponseWriter
	got404 bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
		hrw.got404 = true
	} else {
		hrw.ResponseWriter.WriteHeader(status)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.got404 {
		return len(p), nil
	}

	return hrw.ResponseWriter.Write(p)
}

func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept"), "text/html") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found")
			return
		}

		index, err := files.Open(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)
			return
		}

		fi, err := index.Stat()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)
			return
		}

		w.Header().Set("Content-Type", "text/html: charset=utf-8")
		http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
	}
}

func main() {
	cfg := config.NewConfig()
	player := internal.NewPlayerService(cfg)
	auth := internal.NewAuthService(cfg)
	handler := handler.NewHandler(cfg, auth, player)

	s := &http.ServeMux{}
	httpFS := http.FS(os.DirFS("public"))
	fileServer := http.FileServer(httpFS)
	// serverIndex := serveFileContents("index.html", httpFS)

	// s.Handle("GET /", intercept404(fileServer, serverIndex))
	s.Handle("GET /", fileServer)
	s.HandleFunc("GET /login", handler.HandleLogin)
	s.HandleFunc("GET /refresh", handler.HandleRefresh)
	s.HandleFunc("GET /login/client", handler.HandleClientLogin)
	s.HandleFunc("GET /logout", handler.HandleLogout)
	s.HandleFunc("GET /callback", handler.HandleCallback)
	s.HandleFunc("GET /api/playing", handler.HandlePlaying)

	fmt.Println("Server up and listening on http://localhost:9001")
	if err := http.ListenAndServe(":9001", s); err != nil {
		panic(err)
	}
}
