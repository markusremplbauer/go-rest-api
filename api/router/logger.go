package router

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf("%-10s %-30s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
