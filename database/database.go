package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	Connect := "user=postgres dbname={login} password={password} host=localhost sslmode=disable"
	db, err := sql.Open("postgres", Connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}
