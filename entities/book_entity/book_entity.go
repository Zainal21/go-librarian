package book_entity

import "time"

type Book struct {
	Id          *string
	BookCode    string
	Title       string
	Description string
	Page        int
	AuthorId    string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
