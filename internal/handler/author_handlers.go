package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "BookStore/internal/database"
	models "BookStore/internal/models"
)

type AuthorHandler struct {
	Storage *db.Storage
}

func NewAuthorHandler(storage *db.Storage) *AuthorHandler {
	return &AuthorHandler{Storage: storage}
}

func (h *AuthorHandler) CreateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	var temp struct {
		Name      string `json:"name"`
		BirthDate string `json:"birth_date"`
		Biography string `json:"biography"`
	}
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	var err error
	author.BirthDate, err = time.Parse("2006-01-02", temp.BirthDate)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	author.Name = temp.Name
	author.Biography = temp.Biography
	createdAuthor, err := h.Storage.CreateAuthor(author)
	if err != nil {
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdAuthor)
}

func (h *AuthorHandler) GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := h.Storage.GetAuthors()
	if err != nil {
		http.Error(w, "Failed to retrieve authors", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (h *AuthorHandler) UpdateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("author_id")
	fmt.Println("UTHOR ID", id)
	var updatedAuthor models.Author
	if err := json.NewDecoder(r.Body).Decode(&updatedAuthor); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := h.Storage.UpdateAuthor(r.Context(), id, updatedAuthor)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to update author", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Author successfully updated"})
}

func (h *AuthorHandler) DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.Storage.DeleteAuthor(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Author successfully deleted"})
}
