package author

type AuthorResponse struct {
	Id        *string `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	CreatedAt *string `json:"create_at"`
	UpdatedAt *string `json:"updated_at"`
}

type AuthorRequest struct {
	Id        *string `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       string  `json:"age"`
	CreatedAt *string `json:"create_at"`
	UpdatedAt *string `json:"updated_at"`
}
