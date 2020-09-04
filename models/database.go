package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	// load environment
	e := godotenv.Load()
	if e != nil {
		log.Fatalln(e)
	}

	// set db environment
	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// set postgres db
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUri,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = conn

	// migrate models
	db.Debug().AutoMigrate(&User{}, &Authentication{})
}

func GetDB() *gorm.DB {
	return db
}