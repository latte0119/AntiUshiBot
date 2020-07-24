package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
GetDDB returns a new dynamodb client
*/
func GetDDB() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(nil)

	svc := dynamodb.New(sess)
	return svc
}
