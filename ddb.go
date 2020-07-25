package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
GetDDB returns a new dynamodb client
*/
func GetDDB() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	if svc == nil {
		log.Fatal("db not found")
	}
	return svc
}
