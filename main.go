package main

import (
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
)

func main() {
	Init()

	//	api := GetTwitterAPI()
	db := GetDB()
	defer db.Close()

	var num int
	db.QueryRow("select count(*) from tweet").Scan(&num)
	row := db.QueryRow("select id,text from tweet where id= $1 ", rand.Intn(num))

	var id int
	var text string
	row.Scan(&id, &text)

	fmt.Println(id, text)
	///api.PostTweet(text+" KOREHA TESUTO DESU", nil)
}
