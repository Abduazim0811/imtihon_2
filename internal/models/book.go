package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID              int `json:"id"`
	Title           string    `json:"title"`
	AuthorID        uuid.UUID `json:"author_id"`
	PublicationDate time.Time `json:"publication_date"`
	ISBN            string    `json:"isbn"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
