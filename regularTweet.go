package main

import (
	"errors"

	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getDDB() (*dynamodb.DynamoDB, error) {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	if svc == nil {
		return nil, errors.New("Error : cannot connect to the DB")
	}
	return svc, nil
}

type item struct {
	Tweet string `json:"tweet"`
}

func fetchTweetList() ([]string, error) {
	db, err := getDDB()
	if err != nil {
		return nil, err
	}
	out, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("aub-tweet"),
	})

	if err != nil {
		return nil, err
	}

	arr := make([]string, len(out.Items))
	for k, att := range out.Items {
		itm := item{}
		if err := dynamodbattribute.UnmarshalMap(att, &itm); err != nil {
			return nil, err
		}
		arr[k] = itm.Tweet
	}

	return arr, nil
}

func (aub *AntiUshiBot) RegularTweet() error {
	arr, err := fetchTweetList()
	if err != nil {
		return err
	}

	if len(arr) == 0 {
		return errors.New("Error : tweet list is empty")
	}

	k := rand.Intn(len(arr))

	return aub.PostTweet(arr[k])
}
