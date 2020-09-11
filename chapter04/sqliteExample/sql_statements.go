package main

const CREATE_BOOK_TABLE = "CREATE TABLE IF NOT EXISTS books(id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)"
const INSERT_BOOK = "INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)"
const SELECT_BOOKS = "SELECT id, name, author FROM books"
const UPDATE_BOOK_NAME = "UPDATE BOOKS SET name=? WHERE id=?"
const DELETE_BOOK_BY_ID = "DELETE FROM books WHERE id=?"