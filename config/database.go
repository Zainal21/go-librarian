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

	return db, nil
}
