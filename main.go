package main

import (
	"os"

	"math/rand"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	Init()

	api := GetTwitterAPI()

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()

	var num int
	db.QueryRow("select count(*) from tweet").Scan(&num)
	row := db.QueryRow("select id,text from tweet where id= $1 ", rand.Intn(num))

	var id int
	var text string
	row.Scan(&id, &text)
	api.PostTweet(text, nil)
}
