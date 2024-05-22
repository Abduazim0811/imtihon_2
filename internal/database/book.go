package db

import (
	models "BookStore/internal/models"
	"context"
	"strings"
)

func (s *Storage) CreateBook(book models.Book) (models.Book, error) {
	var insertedBook models.Book
	row := s.db.QueryRow("INSERT INTO books(title, author_id, publication_date, isbn, description) VALUES($1, $2, $3, $4, $5) RETURNING title, author_id, publication_date, isbn, description",
		book.Title, book.AuthorID, book.PublicationDate, book.ISBN, book.Description)

	if err := row.Scan(&insertedBook.Title, &insertedBook.AuthorID, &insertedBook.PublicationDate, &insertedBook.ISBN, &insertedBook.Description); err != nil {
		return models.Book{}, err
	}

	return insertedBook, nil
}
func (s *Storage) GetBooks() ([]models.Book, error) {
	rows, err := s.db.Query("SELECT book_id, title, author_id, publication_date, isbn, description FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublicationDate, &book.ISBN, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
func (s *Storage) GetBookById(id string) (models.Book, error) {
	var book models.Book
	err := s.db.QueryRow("SELECT book_id, title, author_id, publication_date, isbn, description FROM books where book_id =$1", id).Scan(
		&book.ID, &book.Title, &book.AuthorID, &book.PublicationDate, &book.ISBN, &book.Description)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}
func (s *Storage) UpdateBook(ctx context.Context, id int, book models.Book) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := `
        UPDATE books
        SET title = $1, author_id = $2, description = $3
        WHERE book_id = $4
    `

	_, err = tx.ExecContext(ctx, query, strings.ToLower(book.Title), book.AuthorID, book.Description, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (s *Storage) DeleteBook(ctx context.Context, id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := `
        DELETE FROM books
        WHERE book_id = $1
    `

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
