package main

import (
	"go-book-management/config"
	"go-book-management/handler/auth"
	"go-book-management/handler/author"
	"go-book-management/handler/book"
	"go-book-management/middleware"
	"go-book-management/repositories/authors"
	"go-book-management/repositories/books"
	"go-book-management/repositories/users"
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
	UserRepository := users.NewUserRepository(db)
	// routes
	router := mux.NewRouter()

	// auth router
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.Login(w, r, UserRepository)
	}).Methods("POST")

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.Register(w, r, UserRepository)
	}).Methods("POST")

	// Books router
	booksRouter := router.PathPrefix("/books").Subrouter()
	booksRouter.Use(middleware.ProtectedMiddleware)

	booksRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		book.GetBooks(w, r, BookRepository)
	}).Methods("GET")
	booksRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.GetBookById(w, r, BookRepository)
	}).Methods("GET")
	booksRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		book.CreateBook(w, r, BookRepository)
	}).Methods("POST")
	booksRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.UpdateBook(w, r, BookRepository)
	}).Methods("PUT")
	booksRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		book.DeleteBook(w, r, BookRepository)
	}).Methods("DELETE")

	// Authors router
	authorRouter := router.PathPrefix("/authors").Subrouter()
	authorRouter.Use(middleware.ProtectedMiddleware)

	authorRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		author.GetAuthors(w, r, AuthorRepository)
	}).Methods("GET")
	authorRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.GetAuthorById(w, r, AuthorRepository)
	}).Methods("GET")
	authorRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		author.CreateAuthor(w, r, AuthorRepository)
	}).Methods("POST")
	authorRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.UpdateAuthor(w, r, AuthorRepository)
	}).Methods("PUT")
	authorRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		author.DeleteAuthor(w, r, AuthorRepository)
	}).Methods("DELETE")

	// register route
	http.Handle("/", router)
	log.Println("Server listen in port 8080")
	http.ListenAndServe(":8080", nil)
}
