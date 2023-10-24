package book

import (
	"go-book-management/repositories/books"
	"go-book-management/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	books, err := BookRepository.GetAllBooks()

	if err != nil {
		log.Println("error", err.Error())
		utils.JsonResponse(w, nil, "ERROR", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, books, "Books", http.StatusOK)
}

func GetBookById(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.JsonResponse(w, nil, "Id cannot empty", http.StatusBadRequest)
		return
	}

	books, err := BookRepository.FindBookById(id)

	if err != nil {
		log.Println("error", err.Error())
		utils.JsonResponse(w, nil, "ERROR", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, books, "Books Detail", http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Book Management Store", "develop" : true}`))
}

func UpdateBook(w http.ResponseWriter, r *http.Request, BookRepository *books.BookRepository) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Book Management Update", "develop" : true}`))
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

	utils.JsonResponse(w, 1, "Book deleted successfully", http.StatusOK)
}
