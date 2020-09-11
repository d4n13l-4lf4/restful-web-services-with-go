package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id int
	name string
	author string
}

func dbOperations(db *sql.DB) {
	statement, _ := db.Prepare(INSERT_BOOK) // Prepare -> for change ops
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database!")

	rows, _ := db.Query(SELECT_BOOKS) // Query -> for read-only ops
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	statement, _ = db.Prepare(UPDATE_BOOK_NAME)
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")

	statement, _ = db.Prepare(DELETE_BOOK_BY_ID)
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Println(err)
	}

	statement, err := db.Prepare(CREATE_BOOK_TABLE)

	if err  != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}
	statement.Exec()
	dbOperations(db)
}