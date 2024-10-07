package config

import (
	"fmt"
	"log"
	"os"

	"crudapi/helper"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DatabaseConnection() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		host := os.Getenv("DB_HOST")
		if host == "" {
			host = "postgres" // change to localhost if you are using docker local postgres database
		}

		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}

		user := os.Getenv("DB_USER")
		if user == "" {
			user = "postgres"
		}

		password := os.Getenv("DB_PASSWORD")
		if password == "" {
			password = "postgres"
		}

		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "postgres"
		}

		databaseURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
