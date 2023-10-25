package auth

type LoginResponse struct {
	Id        *string `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt *string `json:"create_at"`
	UpdatedAt *string `json:"updated_at"`
	Token     *string `json:"token"`
}
