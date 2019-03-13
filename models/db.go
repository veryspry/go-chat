package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// The postgres database dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Scope this at highest file level so GetDB has access to it
var db *gorm.DB
var dbURI string

func init() {

	// Load the .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// Get vars from environment
	dbUsername := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//Build connection string
	dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUsername, dbName, dbPassword)
	fmt.Println(dbURI)

	// Connect to the database
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	//Database migration
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Ticket{})
	db.Debug().AutoMigrate(&Message{})
	db.Debug().AutoMigrate(&Conversation{})
}

// GetDB returns a reference to the db
func GetDB() *gorm.DB {
	return db
}

// GetDBURI returns the formatted DBURI
func GetDBURI() string {
	return dbURI
}
