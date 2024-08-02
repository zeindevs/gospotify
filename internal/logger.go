package internal

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start).String())
	})
}

func LoggerAsset(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start).String())
	})
}
