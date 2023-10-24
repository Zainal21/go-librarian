package authors

import (
	"database/sql"
	"go-book-management/entities/author_entity"
	"time"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		db,
	}
}

func (r AuthorRepository) GetAllAuthors() ([]author_entity.Author, error) {
	rows, err := r.db.Query("SELECT * FROM authors")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []author_entity.Author

	for rows.Next() {
		var AuthorEntity author_entity.Author
		if err := rows.Scan(
			&AuthorEntity.Id,
			&AuthorEntity.FirstName,
			&AuthorEntity.LastName,
			&AuthorEntity.Age,
			&AuthorEntity.CreatedAt,
			&AuthorEntity.UpdatedAt,
		); err != nil {
			return nil, err
		}
		authors = append(authors, AuthorEntity)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r AuthorRepository) FindAuthorById(id string) (*author_entity.Author, error) {
	query := "SELECT * FROM authors where id = ?"
	row := r.db.QueryRow(query, id)

	var author author_entity.Author

	err := row.Scan(&author.Id, &author.FirstName, &author.LastName, &author.Age, &author.CreatedAt, &author.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &author, err

}

func (r *AuthorRepository) CreateAuthor(author *author_entity.Author) error {
	query := "INSERT INTO authors (id, first_name, last_name, age, created_at, updated_at) VALUES (UUID(),?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, author.FirstName,
		author.LastName,
		author.Age,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepository) UpdateAuthor(payload *author_entity.Author, id string) error {
	query := "UPDATE authors SET first_name = ?, last_name = ?, age = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, payload.FirstName, payload.LastName, payload.Age, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthorRepository) DeleteAuthor(id string) error {
	query := "DELETE FROM authors WHERE id = ?"
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
