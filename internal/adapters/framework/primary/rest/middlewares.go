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
			"%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			r.Response.Status, // TODO: Add response status code. Throws error
		)
		next.ServeHTTP(w, r)
	})
}
