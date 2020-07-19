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

	var num int
	db.QueryRow("select count(*) from tweet").Scan(&num)
	fmt.Println(num)
	row := db.QueryRow("select id,text from tweet where id= $1 ", rand.Intn(3))

	var id int
	var text string
	row.Scan(&id, &text)
	///api.PostTweet(text+" KOREHA TESUTO DESU", nil)
}
