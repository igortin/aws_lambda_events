package main

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewClient initialize client
func NewClient() (*Client, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Client{
		dbsvc: dynamodb.New(sess),
	}, nil
}

// Client struct
type Client struct {
	dbsvc *dynamodb.DynamoDB
}

//CustomPutItem Method
func (client *Client) CustomPutItem(tableName string, e *Event) (*dynamodb.PutItemOutput, error) {
	// Convert string to time.Time.Unix
	timeEvent, _ := time.Parse(time.RFC3339, e.EventTime)

	UnixTime := strconv.FormatInt(timeEvent.Unix(), 10)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ApiID": {
				S: aws.String(e.EventID),
			},
			"Timestamp": {
				S: aws.String(UnixTime),
			},
			"EventName": {
				S: aws.String(e.EventName),
			},
			"RequestParameters": {
				S: aws.String(string(e.RequestParameters)),
			},
			"ResponseElements": {
				S: aws.String(string(e.ResponseElements)),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String(tableName),
	}
	res, err := client.dbsvc.PutItem(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
