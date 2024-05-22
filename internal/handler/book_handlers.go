package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "BookStore/internal/database"
	models "BookStore/internal/models"
)

type BookHandler struct {
	Storage *db.Storage
}

func NewBookHandler(storage *db.Storage) *BookHandler {
	return &BookHandler{Storage: storage}
}

func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	createdBook, err := h.Storage.CreateBook(book)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.Storage.GetBooks()
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	book, err := h.Storage.GetBookById(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = h.Storage.UpdateBook(r.Context(), book.ID, book)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book successfully updated"})
}

func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.Storage.DeleteBook(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book successfully deleted"})
}
