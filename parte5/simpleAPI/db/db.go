package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	con := "user=postgres dbname=alura_loja password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	}
	return db
}