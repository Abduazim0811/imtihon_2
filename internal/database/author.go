package db

import (
	models "BookStore/internal/models"
	"context"
	"log"
	"strings"
)

func (s *Storage) CreateAuthor(author models.Author) (models.Author, error) {
	var insertedAuthor models.Author
	row := s.db.QueryRow("INSERT INTO authors(name, birth_date, biography) VALUES($1, $2, $3) RETURNING name, birth_date, biography",
		author.Name, author.BirthDate, author.Biography)
	if err := row.Scan(&insertedAuthor.Name, &insertedAuthor.BirthDate, &insertedAuthor.Biography); err != nil {
		log.Println(err)
		return models.Author{}, err
	}

	return insertedAuthor, nil
}

func (s *Storage) GetAuthors() ([]models.Author, error) {
	rows, err := s.db.Query("SELECT name, birth_date, biography FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.Name, &author.BirthDate, &author.Biography); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *Storage) UpdateAuthor(ctx context.Context, id string, author models.Author) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := `
        UPDATE authors
        SET name = $1, birth_date = $2, biography = $3
        WHERE author_id = $4
    `

	_, err = tx.ExecContext(ctx, query, strings.ToLower(author.Name), author.BirthDate, author.Biography, id)
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

func (s *Storage) DeleteAuthor(ctx context.Context, id string) error {
    tx, err := s.db.Begin()
    if err != nil {
        return err
    }

    query := `
        DELETE FROM authors
        WHERE author_id = $1
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
