package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type MyEvent struct {
	Name string `json:"name"`
}

/*
	func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
		return fmt.Sprintf("Hello %s!", name.Name), nil
	}
*/

type leaderboardDynamoDBItem struct {
	ActId    string `json:"actid"`
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
	Rank     int    `json:"leaderboardRank"`
	Rating   int    `json:"rankedRating"`
	Wins     int    `json:"numberOfWins"`
}

var activeActId string = "573f53ac-41a5-3a7d-d9ce-d6a6298e5704"
var tablename = "val_leaderboards"

func HandleRequest(ctx context.Context, name MyEvent) ([]leaderboardDynamoDBItem, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	result := findAll(svc, activeActId)
	return result, nil
}

func findAll(svc dynamodbiface.DynamoDBAPI, actid string) []leaderboardDynamoDBItem {
	filt := expression.Name("actid").Equal(expression.Value(actid))
	proj := expression.NamesList(expression.Name("gameName"), expression.Name("leaderboardRank"), expression.Name("numberOfWins"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tablename),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	itemList := []leaderboardDynamoDBItem{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &itemList)
	if err != nil {
		log.Fatalf("Got error unmarshalling: %s", err)
	}

	return itemList
}

func main() {
	lambda.Start(HandleRequest)
}
