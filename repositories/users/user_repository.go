package users

import (
	"database/sql"
	"go-book-management/entities/user_entity"
	"go-book-management/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// user repository
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

type UserLoginResponse struct {
	ID        string
	Username  string
	Email     string
	CreatedAt string
	UpdatedAt string
}

func (r UserRepository) Register(payload *user_entity.User) error {

	hashPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		return err
	}

	query := "INSERT INTO users (id, username, password, email, created_at) VALUES (UUID(),?, ?, ?, ?)"
	_, err = r.db.Exec(query, payload.Username, hashPassword, payload.Email, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) VerifyPassword(username, password string) bool {
	storedHashedPassword, err := r.getHashedPassword(username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func (r *UserRepository) getHashedPassword(username string) (string, error) {
	query := "SELECT password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var storedHashedPassword string
	err := row.Scan(&storedHashedPassword)
	if err != nil {
		return "", err
	}

	return storedHashedPassword, nil
}

func (r *UserRepository) FindUserByUserName(username string) (*user_entity.User, error) {
	query := "SELECT id, username, password, email, created_at, updated_at FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user user_entity.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
