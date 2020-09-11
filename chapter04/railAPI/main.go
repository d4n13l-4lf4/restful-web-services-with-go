package main

import (
	"database/sql"
	"log"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter04/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(db)
}