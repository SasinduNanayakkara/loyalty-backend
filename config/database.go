package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() {

	url := os.Getenv("DB_URL")
	
	db, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	} else {
		log.Println("Database connection established successfully", db)
	}
}

