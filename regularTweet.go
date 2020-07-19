package main

import (
	"log"
	"math/rand"
)

/*
RegularTweet tweets regularly
*/
func RegularTweet() {

	api := GetTwitterAPI()
	db := GetDB()
	defer db.Close()

	var num int
	db.QueryRow("select count(*) from tweet").Scan(&num)
	row := db.QueryRow("select id,text from tweet where id= $1 ", rand.Intn(num))

	var id int
	var text string
	row.Scan(&id, &text)

	log.Println(id, text)
	api.PostTweet(text, nil)
}
