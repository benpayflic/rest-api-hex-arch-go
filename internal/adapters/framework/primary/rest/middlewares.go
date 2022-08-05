package rest

import (
	"log"
	"net/http"
)

// Set response content header
func responseManagerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// log all incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s\t%s",
			r.Method,
			r.RequestURI,
		)
		next.ServeHTTP(w, r)
	})
}
