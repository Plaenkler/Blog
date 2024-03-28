package server

import (
	"net/http"
)

func controlCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000")
		next.ServeHTTP(w, r)
	})
}
