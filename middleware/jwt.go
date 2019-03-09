package middleware

import (
	"context"
	"fmt"
	"go-auth/models"
	u "go-auth/utils"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//List of endpoints that don't require auth
		notAuth := []string{"/login", "/user/new", "ws"}
		//current request path
		requestPath := r.URL.Path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		//Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")

		//Token is missing, returns with error code 403 Unauthorized
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//The token normally comes in format "Bearer {token-body}", we check if the retrieved token matched this requirement
		split := strings.Split(tokenHeader, " ")
		if len(split) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Grab the token part, what we are interested in
		tokenPart := split[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//Malformed token, returns with http code 403 as usual
		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Token is invalid, maybe not signed on this server
		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		usrID := fmt.Sprintf("UserID %d", tk.UserID) //Useful for monitoring
		fmt.Println(usrID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)

		fmt.Print("ctx", ctx)
		r = r.WithContext(ctx)
		//proceed in the middleware chain
		next.ServeHTTP(w, r)
	})
}
