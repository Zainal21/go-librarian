package book

import (
	"encoding/json"
	"errors"
	"go-book-management/entities/book_entity"
	"go-book-management/repositories/books"
	"go-book-management/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookPayload struct {
	Title       string
	Description string
	Page        int
	AuthorId    string
}

const (
	ErrRequiredFields     = "All fields are required"
	ErrInvalidAgeFormat   = "Invalid age format"
	ErrBookIDEmpty        = "ID cannot be empty"
	ErrBookCreationFailed = "Book creation failed"
	ErrBookUpdateFailed   = "Book update failed"
)

func parseBookPayload(r *http.Request) (BookPayload, error) {
	var request BookRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return BookPayload{}, errors.New(ErrBookCreationFailed)
	}

	title := request.Title
	author_id := request.AuthorId
	description := request.Description
	page := request.Page

	if title == "" || author_id == "" || description == "" || page == "" {
		return BookPayload{}, errors.New(ErrRequiredFields)
	}

	pageInt, err := parseInt(page)
	if err != nil {
		return BookPayload{}, errors.New(ErrInvalidAgeFormat)
	}

	return BookPayload{
		Title:       title,
		Description: description,
		Page:        pageInt,
		AuthorId:    author_id,
	}, nil
}

func parseInt(value string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func GetBooks(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	books, err := BookRepository.GetAllBooks()

	if err != nil {
		log.Println("error", err.Error())
		utils.JsonResponse(w, nil, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, books, "SUCCESS", http.StatusOK)
}

func GetBookById(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	books, err := BookRepository.FindBookById(id)
	log.Println("books : ", books)
	if err != nil {
		log.Println("error", err.Error())
		utils.JsonResponse(w, nil, "NOT FOUND", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, books, "SUCCESS", http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	data, err := parseBookPayload(r)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// payload
	var payload book_entity.Book
	payload.Title = data.Title
	payload.Description = data.Description
	payload.Page = data.Page
	payload.AuthorId = data.AuthorId

	err = BookRepository.Create(&payload)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}
	utils.JsonResponse(w, 1, "SUCCESS", http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	data, err := parseBookPayload(r)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// payload
	var payload book_entity.Book
	payload.Title = data.Title
	payload.Description = data.Description
	payload.Page = data.Page
	payload.AuthorId = data.AuthorId

	err = BookRepository.Update(&payload, id)

	if err != nil {
		utils.JsonResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JsonResponse(w, 1, "SUCCESS", http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	err := BookRepository.DeleteBook(id)

	if err != nil {
		log.Println("error when delete", err.Error())
		utils.JsonResponse(w, nil, "Data not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, 1, "SUCCESS", http.StatusOK)
}
