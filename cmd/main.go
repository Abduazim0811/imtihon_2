package main

import (
	"log"
	"net/http"

	database "BookStore/internal/database"
	"BookStore/internal/handler"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres driver
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db, e := database.ConnectDB()
	if e != nil {
		log.Println(e)
		log.Fatal("---------DB connection could not be set --------------------")
	}

	storage := database.NewStorage(db)
	bookHandler := handler.NewBookHandler(storage)
	authorHandler := handler.NewAuthorHandler(storage)

	http.HandleFunc("/books", bookHandler.CreateBookHandler)
	http.HandleFunc("/books/", bookHandler.GetBooksHandler)
	http.HandleFunc("/update-book", bookHandler.UpdateBookHandler)
	http.HandleFunc("/delete-book", bookHandler.DeleteBookHandler)

	http.HandleFunc("/author", authorHandler.CreateAuthorHandler)
	http.HandleFunc("/authors", authorHandler.GetAuthorsHandler)
	http.HandleFunc("/update-author", authorHandler.UpdateAuthorHandler)
	http.HandleFunc("/delete-author", authorHandler.DeleteAuthorHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
