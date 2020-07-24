package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type item struct {
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
	itm := item{}
	dynamodbattribute.UnmarshalMap(arr[k], &itm)

	fmt.Println(itm.Tweet)

	api := GetTwitterAPI()
	api.PostTweet(itm.Tweet, nil)
}
