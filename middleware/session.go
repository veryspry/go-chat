package middleware

import (
	"fmt"
	"go-auth/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/joho/godotenv"
)

func SessionMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Load the .env file
		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
		// Get token_secret
		tokenSecret := os.Getenv("token_secret")
		// Get db uri
		dbURI := models.GetDBURI()

		// Fetch new store.
		store, err := pgstore.NewPGStore(dbURI, []byte(tokenSecret))
		if err != nil {
			log.Fatalf(err.Error())
		}

		defer store.Close()

		// Run a background goroutine to clean up expired sessions from the database.
		defer store.StopCleanup(store.Cleanup(time.Minute * 5))

		// Get a session.
		session, err := store.Get(r, "session-key")
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Add a value.
		session.Values["foo"] = "bar"

		// Save.
		if err = session.Save(r, w); err != nil {
			log.Fatalf("Error saving session: %v", err)
		}

		// Delete session.
		session.Options.MaxAge = -1
		if err = session.Save(r, w); err != nil {
			log.Fatalf("Error saving session: %v", err)
		}
	})
}
