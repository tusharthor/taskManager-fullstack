package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting db: %v", err)
	}

	//ping
	if err := DB.Ping(); err != nil {
		log.Fatal("database connection is not alive")
	}

	log.Println("connected to db")

}
