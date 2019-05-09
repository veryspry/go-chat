package middleware

import (
	"fmt"
	"net/http"
)

// LogReqBody Logs out an http Request Body
func LogReqBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v Request -- REQUEST BODY:\n%v", r.Method, r.Body)
	})
}
