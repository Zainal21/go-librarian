package books

import (
	"database/sql"
	"fmt"
	"go-book-management/entities/book_entity"
	"log"
	"strconv"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

type Book struct {
	Id          *string `json:"id"`
	BookCode    string  `json:"book_code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Page        int     `json:"page"`
	AuthorName  string  `json:"author_name"`
	AuthorId    string  `json:"author_id"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db,
	}
}

func (r BookRepository) GetAllBooks() ([]Book, error) {
	query := `SELECT
        books.id,
        books.book_code,
        books.title,
        books.description,
        books.page,
        books.author_id,
		books.created_at,
		books.updated_at,
        CONCAT(authors.first_name, ' ', authors.last_name) as author_name
    FROM
        books
    JOIN authors ON
        books.author_id = authors.id`

	log.Println("Book selected : ", query)
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		var authorName string // Define a variable for author name
		if err := rows.Scan(
			&book.Id,
			&book.BookCode,
			&book.Title,
			&book.Description,
			&book.Page,
			&book.AuthorId,
			&authorName, // Scan author name into a variable
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return nil, err
		}
		book.AuthorName = authorName // Assign the author name to the struct field
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r BookRepository) FindBookById(id string) (*Book, error) {
	query := `SELECT
			books.id,
			books.book_code,
			books.title,
			books.description,
			books.page,
			books.author_id,
			books.created_at,
			books.updated_at,
			CONCAT(authors.first_name, ' ', authors.last_name) as author_name
		FROM
			books
		JOIN authors ON
			books.author_id = authors.id
		WHERE books.id = ?`
	log.Println("Book find by id : ", query)

	row := r.db.QueryRow(query, id)

	var book Book

	err := row.Scan(&book.Id,
		&book.BookCode,
		&book.Title,
		&book.Description,
		&book.Page,
		&book.AuthorId,
		&book.AuthorName,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &book, err
}

func (r *BookRepository) Create(payload *book_entity.Book) error {
	query := `
		INSERT INTO books 
			(id, book_code, title, description, page, author_id, created_at, updated_at) 
		VALUES 
			(UUID(), ?, ? , ? , ? , ? , ?, ?)`

	bookCode, err := r.GenerateUniqueBookCode()

	if err != nil {
		log.Println("Generate Book Code : ", err)
		return err
	}

	_, err = r.db.Exec(query, bookCode, payload.Title, payload.Description, payload.Page, payload.AuthorId, time.Now(), time.Now())
	if err != nil {
		log.Println("Execute query : ", err)
		return err
	}

	return nil
}

func (r *BookRepository) Update(payload *book_entity.Book, id string) error {
	query := `
        UPDATE books
        SET
            title = ?,
            description = ?,
            page = ?,
            author_id = ?,
            updated_at = ?
        WHERE
            id = ?
    `

	_, err := r.db.Exec(query,
		payload.Title,
		payload.Description,
		payload.Page,
		payload.AuthorId,
		time.Now(),
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) GenerateUniqueBookCode() (string, error) {
	// Get the current month and year
	currentTime := time.Now()
	month := currentTime.Format("01")
	year := currentTime.Format("2006")
	yearFormat := currentTime.Format("06")

	// Initialize the baseOrdered with "0000"
	lastOrder := "0000"

	// Query the database to find the highest value of the last 4 characters in book_code
	query := `
        SELECT COALESCE(MAX(CAST(RIGHT(book_code, 4) AS SIGNED)), 0000) AS last_order
        FROM books
        WHERE MONTH(created_at) = ? AND YEAR(created_at) = ?
    `
	err := r.db.QueryRow(query, month, year).Scan(&lastOrder)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	lastOrderInt, err := strconv.Atoi(lastOrder)

	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error:", err)
		return "", err
	}

	// Calculate the next ordered value
	nextOrdered := lastOrderInt + 1
	// Generate the unique code
	uniqueCode := fmt.Sprintf("BOOK%s%s%04d", month, yearFormat, nextOrdered)
	return uniqueCode, nil
}

func (r *BookRepository) DeleteBook(id string) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
