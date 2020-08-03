package main

import (
	"context"
	"encoding/json"

	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	awsRegion  = "us-east-1"
	awsProfile = "private"
	tableName  = "ApiEvents"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {

	e := NewEvent()
	err := json.Unmarshal([]byte(event.Detail), e)
	if err != nil {
		log.Println(err)
	}
	client, err := NewClient()
	if err != nil {
		log.Println(err)
	}
	switch e.EventName {
	case "StartInstances":
		_, err := client.CustomPutItem(tableName, e)
		if err != nil {
			log.Println(err)
		}
	case "StopInstances":

		_, err := client.CustomPutItem(tableName, e)

		if err != nil {
			log.Println(err)
		}

	case "CreateUser":

		_, err := client.CustomPutItem(tableName, e)
		if err != nil {
			log.Println(err)
		}
	case "DeleteUser":

		_, err := client.CustomPutItem(tableName, e)
		if err != nil {
			log.Println(err)
		}
	default:
		fmt.Println("No event logic")
	}
}

func main() {
	lambda.Start(handler)
}
