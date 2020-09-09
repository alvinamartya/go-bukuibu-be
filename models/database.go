package models

import (
	"fmt"
	"github.com/alvinamartya/go-bukuibu-be/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	// load environments
	dbUser, err := utils.GetEnvVar("db_user")
	dbPass, err := utils.GetEnvVar("db_pass")
	dbName, err := utils.GetEnvVar("db_name")
	dbHost, err := utils.GetEnvVar("db_host")
	dbPort, err := utils.GetEnvVar("db_port")
	if err != nil {
		log.Fatalln(err)
	}

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
	db.Debug().AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	return db
}
