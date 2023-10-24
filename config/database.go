package config

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func DbConnection() (*sql.DB, error) {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Println("err : " + err.Error())
		return nil, err
	}

	// create migration table authors
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS authors (
	// 		id char(36) PRIMARY KEY,
	// 		first_name VARCHAR(50) NOT NULL,
	// 		last_name VARCHAR(64) NOT NULL,
	// 		age INT NOT NULL DEFAULT 1,
	// 		created_at TIMESTAMP  NULL,
	// 		updated_at TIMESTAMP  NULL
	// 	);
	// `)

	// // create migration table books
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
	// 		id char(36) PRIMARY KEY,
	// 		book_code VARCHAR(50) NOT NULL,
	// 		title VARCHAR(64) NOT NULL,
	// 		description VARCHAR(64) NOT NULL,
	// 		page INT NOT NULL DEFAULT 1,
	// 		author_id INT,
	// 		created_at TIMESTAMP  NULL,
	// 		updated_at TIMESTAMP  NULL
	// 	);
	// `)

	// if err != nil {
	// 	log.Println("err : " + err.Error())
	// 	return nil, err
	// }

	// log.Println("Database migration completed successfully.")

	// seeder database

	// Sample user data
	// authors := []struct {
	// 	first_name string
	// 	last_name  string
	// 	age        int
	// }{
	// 	{"Muhamad Zainal", "Arifin", 20},
	// 	{"Feri", "Irwandi", 20},
	// 	{"Arifin", "Husein", 20},
	// }

	// // Insert sample user data
	// insertAuthor := "INSERT INTO authors (id, first_name, last_name, age) VALUES (UUID(),?, ?, ?)"

	// for _, author := range authors {
	// 	_, err := db.Exec(insertAuthor, author.first_name, author.last_name, author.age)
	// 	if err != nil {
	// 		fmt.Println("Error inserting author:", err)
	// 		return DB, err
	// 	}
	// 	fmt.Printf("author %s inserted successfully.\n", author.first_name)
	// }
	// log.Println("Seeder completed successfully.")
	return db, nil
}
