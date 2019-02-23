package queries

import (
	"errors"
	"fmt"
	"go-auth/models"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// CreateUser creates a new user
func CreateUser(email, password string) error {

	e := godotenv.Load() // Load the .env file
	if e != nil {
		fmt.Print(e)
	}

	dbUsername := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUsername, dbName, dbPassword)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return errors.New("Failed to connect to the database")
	}

	defer db.Close()

	// Migrate the schema
	db.Debug().AutoMigrate(&models.User{})

	user := models.User{Email: email, Password: password}
	// res := fmt.Sprintf("email=%s password=%s", user.Email, user.Password)
	// fmt.Println(res)

	// Create an item
	res := db.Create(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
