package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

/*
GetDB gets the DB
*/
func GetDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer db.Close()
	return db
}
