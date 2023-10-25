package book

type BookRequest struct {
	Title       string `json:"title"`
	AuthorId    string `json:"author_id"`
	Description string `json:"description"`
	Page        string `json:"page"`
}
