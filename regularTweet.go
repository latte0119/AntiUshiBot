package main

import (
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Tweet string `json:"tweet"`
}

/*
RegularTweet tweets regularly
*/
func RegularTweet() {
	db := GetDDB()
	out, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("aub-tweet"),
	})

	if err != nil {
		log.Fatal(err)
	}

	arr := out.Items

	k := rand.Intn(len(arr))
	item := Item{}
	dynamodbattribute.UnmarshalMap(arr[k], &item)

	log.Println(item.Tweet)

	api := GetTwitterAPI()
	api.PostTweet(item.Tweet, nil)
}
