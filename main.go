package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db := SetupDB()
	rows, err := db.Query("SELECT title, author, description FROM books")

	PanicIf(err)
	defer rows.Close()

	var title, author, description string
	for rows.Next() {
		err := rows.Scan(&title, &author, &description)
		PanicIf(err)
		fmt.Println("Title: %s\nAuthor: %s\nDescription: %s\n\n", title, author, description)
	}
}
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// dials the database, returning any error
func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=dev password=password dbname=test sslmode=disable")
	PanicIf(err)

	return db
}
