package main

import (
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Tweet string `json:"tweet"`
}

func RegularTweet() error {
	db, err := GetDDB()
	if err != nil {
		return err
	}
	out, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("aub-tweet"),
	})

	if err != nil {
		return err
	}

	arr := out.Items

	k := rand.Intn(len(arr))
	item := Item{}
	if err := dynamodbattribute.UnmarshalMap(arr[k], &item); err != nil {
		return err
	}

	api := GetTwitterAPI()
	if _, err := api.PostTweet(item.Tweet, nil); err != nil {
		return err
	}
	return nil
}
