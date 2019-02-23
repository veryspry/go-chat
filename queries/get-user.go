package queries

import (
	"errors"
	"fmt"
	"go-auth/models"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// GetUser Get a single user
func GetUser(email string) (*models.User, error) {

	e := godotenv.Load() // Load the .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to the database")
	}

	defer db.Close()

	user := &models.User{}

	// Find the user
	db.Table("users").Where("email = ?", email).First(user)
	if user.Email == "" {
		return user, errors.New("User not found!\n")
	}

	return user, nil
}
