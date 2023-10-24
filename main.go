package main

import (
	"go-book-management/config"
	"go-book-management/handler/author"
	"go-book-management/handler/book"
	"go-book-management/repositories/authors"
	"go-book-management/repositories/books"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loadded environment variable")
	}

	// db connection
	db, err := config.DbConnection()

	if err != nil {
		log.Fatal("error connect to database")
	}
	// initialize repository
	AuthorRepository := authors.NewAuthorRepository(db)
	BookRepository := books.NewBookRepository(db)
	// routes
	router := mux.NewRouter()

	// Books
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		book.GetBooks(w, r, BookRepository)
	}).Methods("GET")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.GetBookById(w, r, BookRepository)
	}).Methods("GET")
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		book.CreateBook(w, r, BookRepository)
	}).Methods("POST")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.UpdateBook(w, r, BookRepository)
	}).Methods("PUT")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.DeleteBook(w, r, BookRepository)
	}).Methods("DELETE")

	// Authors
	router.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		author.GetAuthors(w, r, AuthorRepository)
	}).Methods("GET")
	router.HandleFunc("/authors/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.GetAuthorById(w, r, AuthorRepository)
	}).Methods("GET")
	router.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		author.CreateAuthor(w, r, AuthorRepository)
	}).Methods("POST")
	router.HandleFunc("/authors/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.UpdateAuthor(w, r, AuthorRepository)
	}).Methods("PUT")
	router.HandleFunc("/authors/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.DeleteAuthor(w, r, AuthorRepository)
	}).Methods("DELETE")

	// register route
	http.Handle("/", router)
	log.Println("Server listen in port 8080")
	http.ListenAndServe(":8080", nil)
}
