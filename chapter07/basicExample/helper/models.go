package helper

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func prepareDBVars() func () (*sql.DB, error) {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "chapter07/basicExample/.env")
	_ = godotenv.Load(path)

	host, port := os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT")
	user, password := os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_PWD")
	dbname := os.Getenv("POSTGRESQL_DB")

	return func () (*sql.DB, error) {
		var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		var err error
		db, err := sql.Open("postgres", connectionString)

		if err != nil {
			return nil, err
		}

		stmt, err := db.Prepare(CREATE_TABLE_WEB_URL)
		if err != nil {
			return nil, err
		}

		_, err = stmt.Exec()

		if err != nil {
			return nil, err
		}

		return db, nil


	}
}

var InitDB = prepareDBVars()

