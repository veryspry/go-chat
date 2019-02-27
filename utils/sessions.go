package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func CreateSession() {
	// Load the .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}
