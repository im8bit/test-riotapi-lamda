package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/im8bit/test-riotapi-library/aws"
	"github.com/im8bit/test-riotapi-library/riot"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) ([]aws.LeaderboardDynamoDBItem, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	var activeActId string = riot.GetActiveActId()

	result := aws.FindAll(svc, activeActId)

	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
