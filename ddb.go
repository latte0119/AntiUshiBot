package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDDB() (*dynamodb.DynamoDB, error) {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	if svc == nil {
		return nil, errors.New("Error : cannot connect to the DB")
	}
	return svc, nil
}
