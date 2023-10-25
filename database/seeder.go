package database

import (
	"fmt"
	"go-book-management/config"
	"log"
)

func init() {
	db, err := config.DbConnection()

	if err != nil {
		log.Println("err : " + err.Error())
	}

	authors := []struct {
		first_name string
		last_name  string
		age        int
	}{
		{"Muhamad Zainal", "Arifin", 20},
		{"Feri", "Irwandi", 20},
		{"Arifin", "Husein", 20},
	}

	// Insert sample user data
	insertAuthor := "INSERT INTO authors (id, first_name, last_name, age) VALUES (UUID(),?, ?, ?)"

	for _, author := range authors {
		_, err := db.Exec(insertAuthor, author.first_name, author.last_name, author.age)
		if err != nil {
			fmt.Println("Error inserting author:", err)
		}
		fmt.Printf("author %s inserted successfully.\n", author.first_name)
	}
	log.Println("Seeder completed successfully.")
}
