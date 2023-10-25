package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id        *string `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt *string `json:"create_at"`
	UpdatedAt *string `json:"updated_at"`
	Token     *string `json:"token"`
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
