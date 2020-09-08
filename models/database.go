package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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
	dbPort := os.Getenv("db_port")

	// set postgres db
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dbUri,
		DefaultStringSize: 256,
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
