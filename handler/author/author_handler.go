package author

import (
	"encoding/json"
	"errors"
	"go-book-management/entities/author_entity"
	"go-book-management/repositories/authors"
	"go-book-management/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AuthorPayload struct {
	FirstName string
	LastName  string
	Age       int
}

const (
	ErrRequiredFields       = "All fields are required"
	ErrInvalidAgeFormat     = "Invalid age format"
	ErrAuthorIDEmpty        = "ID cannot be empty"
	ErrAuthorCreationFailed = "Author creation failed"
	ErrAuthorUpdateFailed   = "Author update failed"
)

func parseInt(value string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func parseAuthorPayload(r *http.Request) (AuthorPayload, error) {
	var request AuthorRequest
	firstName := request.FirstName
	lastName := request.LastName
	ageStr := request.Age

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return AuthorPayload{}, errors.New(ErrAuthorCreationFailed)
	}

	if firstName == "" || lastName == "" || ageStr == "" {
		return AuthorPayload{}, errors.New(ErrRequiredFields)
	}

	age, err := parseInt(ageStr)
	if err != nil {
		return AuthorPayload{}, errors.New(ErrInvalidAgeFormat)
	}

	return AuthorPayload{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}, nil
}

func GetAuthors(w http.ResponseWriter, r *http.Request, AuthorRepository *authors.AuthorRepository) {
	// get all authors
	authors, err := AuthorRepository.GetAllAuthors()

	log.Println("authors", authors)
	if err != nil {
		log.Println("error", err.Error())
		utils.JsonResponse(w, nil, "ERROR", http.StatusInternalServerError)
		return
	}

	var authorResponses []AuthorResponse
	for _, author := range authors {
		authorResponses = append(authorResponses, MapAuthorToResponse(&author))
	}

	utils.JsonResponse(w, authorResponses, "get all authors", http.StatusOK)
}

func MapAuthorToResponse(author *author_entity.Author) AuthorResponse {
	return AuthorResponse{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		Age:       author.Age,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func GetAuthorById(w http.ResponseWriter, r *http.Request, AuthorRepository *authors.AuthorRepository) {
	vars := mux.Vars(r)
	id := vars["id"]

	authors, err := AuthorRepository.FindAuthorById(id)

	if err != nil {
		utils.JsonResponse(w, nil, "Data not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, AuthorResponse{
		Id:        authors.Id,
		FirstName: authors.FirstName,
		LastName:  authors.LastName,
		Age:       authors.Age,
		CreatedAt: authors.Id,
		UpdatedAt: authors.Id,
	}, "Detail Author", http.StatusOK)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request, AuthorRepository *authors.AuthorRepository) {
	data, err := parseAuthorPayload(r)
	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var payload author_entity.Author
	payload.FirstName = data.FirstName
	payload.LastName = data.LastName
	payload.Age = data.Age

	err = AuthorRepository.CreateAuthor(&payload)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JsonResponse(w, 1, "New author created successfully", http.StatusCreated)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request, AuthorRepository *authors.AuthorRepository) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	data, err := parseAuthorPayload(r)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var payload author_entity.Author
	payload.FirstName = data.FirstName
	payload.LastName = data.LastName
	payload.Age = data.Age

	err = AuthorRepository.UpdateAuthor(&payload, id)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JsonResponse(w, 1, "Author updated successfully", http.StatusOK)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request, AuthorRepository *authors.AuthorRepository) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	err := AuthorRepository.DeleteAuthor(id)

	if err != nil {
		log.Println("error when delete", err.Error())
		utils.JsonResponse(w, nil, "Data not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, 1, "Author deleted successfully", http.StatusOK)
}
