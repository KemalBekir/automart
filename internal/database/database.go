package database

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB() {
	connStr := "user=postgres dbname=go_cms sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database.")
}
