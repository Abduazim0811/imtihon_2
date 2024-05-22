package main

import (
	"log"
	"net/http"

	database "BookStore/internal/database"
	"BookStore/internal/handler"
	"BookStore/internal/middleware"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres driver
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		log.Fatal("DB connection could not be set")
	}

	storage := database.NewStorage(db)
	bookHandler := handler.NewBookHandler(storage)
	authorHandler := handler.NewAuthorHandler(storage)

	http.HandleFunc("/books", middleware.ApplyMiddleware(bookHandler.CreateBookHandler, middleware.Logging,middleware.Logging))
	http.HandleFunc("/books/", middleware.ApplyMiddleware(bookHandler.GetBooksHandler,middleware.Logging,middleware.ErrorHandling))
	http.HandleFunc("/update-book", middleware.ApplyMiddleware(bookHandler.UpdateBookHandler,middleware.Logging,middleware.ErrorHandling))
	http.HandleFunc("/delete-book", middleware.ApplyMiddleware(bookHandler.DeleteBookHandler,middleware.Logging,middleware.ErrorHandling))

	http.HandleFunc("/author", middleware.ApplyMiddleware(authorHandler.CreateAuthorHandler,middleware.Logging,middleware.ErrorHandling))
	http.HandleFunc("/authors", middleware.ApplyMiddleware(authorHandler.GetAuthorsHandler,middleware.Logging,middleware.ErrorHandling))
	http.HandleFunc("/update-author", middleware.ApplyMiddleware(authorHandler.UpdateAuthorHandler,middleware.Logging,middleware.ErrorHandling))
	http.HandleFunc("/delete-author", middleware.ApplyMiddleware(authorHandler.DeleteAuthorHandler,middleware.Logging,middleware.ErrorHandling))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
