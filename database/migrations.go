package database

import (
	"go-book-management/config"
	"log"
)

func CreateTable() {
	db, err := config.DbConnection()

	if err != nil {
		log.Println("err : " + err.Error())
	}
	// create migration table authors
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS authors (
			id char(36) PRIMARY KEY,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(64) NOT NULL,
			age INT NOT NULL DEFAULT 1,
			created_at TIMESTAMP  NULL,
			updated_at TIMESTAMP  NULL
		);
	`)

	if err != nil {
		log.Println("err : " + err.Error())
	}

	// create migration table books
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
			id char(36) PRIMARY KEY,
			book_code VARCHAR(50) NOT NULL,
			title VARCHAR(64) NOT NULL,
			description VARCHAR(64) NOT NULL,
			page INT NOT NULL DEFAULT 1,
			author_id char(36,
			created_at TIMESTAMP  NULL,
			updated_at TIMESTAMP  NULL
		);
	`)

	if err != nil {
		log.Println("err : " + err.Error())
	}

	log.Println("Database migration completed successfully.")
}
