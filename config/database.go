package config

import (
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	url := os.Getenv("DB_URL")
	var err error

	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Database connection established successfully")

}
