package book_entity

type Book struct {
	Id          *string
	BookCode    string
	Title       string
	Description string
	Page        int
	AuthorId    string
	CreatedAt   *string
	UpdatedAt   *string
}
