package model

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Biography string    `json:"biography"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}