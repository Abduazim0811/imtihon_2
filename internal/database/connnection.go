package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// storage "BookStore/internal/storage"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		return nil, err
	}

	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	// fmt.Println(host)

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}

	if err := psqlConn.Ping(); err != nil {
		log.Println("error while connecting to the database")
		return nil, err
	}
	log.Println("connection to the database has been created")
	return psqlConn, nil

}
