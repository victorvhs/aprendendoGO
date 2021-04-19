package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaDataBase() *sql.DB {
	con := "user=postgres dbname=simpleAPI password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	}
	return db
}