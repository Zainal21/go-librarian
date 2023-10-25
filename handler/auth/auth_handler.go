package auth

import (
	"encoding/json"
	"errors"
	"go-book-management/entities/user_entity"
	"go-book-management/repositories/users"
	"go-book-management/utils"
	"log"
	"net/http"
)

type UserRegisterPayload struct {
	Username string
	Password string
	Email    string
}

const (
	ErrRequiredFields     = "All fields are required"
	ErrPasswordNotMatch   = "Password & Confirm password doesn't match"
	ErrInvalidAgeFormat   = "Invalid age format"
	ErrUserIDEmpty        = "ID cannot be empty"
	ErrUserCreationFailed = "User register failed"
	ErrUserUpdateFailed   = "User update failed"
)

func parseUserRegisterPayload(r *http.Request) (UserRegisterPayload, error) {
	var request RegisterRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return UserRegisterPayload{}, errors.New(ErrUserCreationFailed)
	}

	username := request.Username
	email := request.Email
	password := request.Password
	passwordConfirm := request.PasswordConfirm

	if username == "" || password == "" || email == "" {
		return UserRegisterPayload{}, errors.New(ErrRequiredFields)
	}

	if password != passwordConfirm {
		return UserRegisterPayload{}, errors.New(ErrPasswordNotMatch)
	}

	return UserRegisterPayload{
		Username: username,
		Password: password,
		Email:    email,
	}, nil

}

func Register(w http.ResponseWriter, r *http.Request, UserRepository *users.UserRepository) {
	data, err := parseUserRegisterPayload(r)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var payload user_entity.User
	payload.Email = data.Email
	payload.Password = data.Password
	payload.Username = data.Username

	err = UserRepository.Register(&payload)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JsonResponse(w, 1, "New user register successfully", http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request, UserRepository *users.UserRepository) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// fetch user from repository
	user, err := UserRepository.FindUserByUserName(username)

	if err != nil || !UserRepository.VerifyPassword(username, password) {
		log.Println(err.Error())
		utils.JsonResponse(w, nil, "Wrong Password", http.StatusUnprocessableEntity)
		return
	}

	// generate token
	token, err := utils.GenerateToken(username)

	if err != nil {
		utils.JsonResponse(w, nil, "Error Generate token", http.StatusInternalServerError)
		return
	}

	if token == "" {
		utils.JsonResponse(w, nil, "Generate Token  is empty", http.StatusInternalServerError)
		return
	}

	var response = LoginResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     &token,
	}

	utils.JsonResponse(w, response, "Login success", http.StatusCreated)
}
