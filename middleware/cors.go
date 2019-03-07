package middleware

import (
	"net/http"
)

// CORSHandler CORS middleware
func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "*")
		headers.Add("Access-Control-Allow-Credentials", "true")
		headers.Add("Content-Type", "application/json")
		headers.Add("Vary", "Origin")
		headers.Add("Vary", "Access-Control-Request-Method")
		headers.Add("Vary", "Access-Control-Request-Headers")
		headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, authorization")
		headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")

		// PREFLIGHT STUFF....
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// If not an OPTIONS request, continue as normal
		next.ServeHTTP(w, r)
	})
}
