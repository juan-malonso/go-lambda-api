package dynamo_events

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/go-lambda-api/apps/libraries/env"
)

// Create Aws Session
var sess, sessErr = session.NewSession(&aws.Config{
	Endpoint:    aws.String(env.GetVar("AWS_DYNAMO_ENDPOINT", "")),
})

// Create Dynamo Client
var dynamoEventsTable = env.GetVar("AWS_DYNAMO_EVENTS_TABLE", "")
var dynamoClient = dynamodb.New(sess)

// Event Structure

type Event struct {
	TransactionId string `json:"transactionId"`
	Event EventItem `json:"event"`
}

type EventItem struct {
	EventId string `json:"eventId"`
	Payload interface{} `json:"payload"`
}

func Marshal(event interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(event)
}


func GetItems(transactionId string) ([]Event, error) {

	// Set Aws Command
	input := &dynamodb.QueryInput{
		TableName: aws.String(dynamoEventsTable),
		KeyConditionExpression: aws.String("id = :tid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":tid": &dynamodb.AttributeValue{S: aws.String(transactionId)},
		},
	}

	// Get Items
	result, err := dynamoClient.Query(input)
	if err != nil { return nil, err }

	// Format Items
	var events []Event
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &events)

	return events, err
}


func PutItem(event Event) error {
	eventItem, err := dynamodbattribute.MarshalMap(event.Event)

	// Set Aws Command
	input := &dynamodb.PutItemInput{
		TableName: aws.String(dynamoEventsTable),
		Item: map[string]*dynamodb.AttributeValue{
			"id": { S: aws.String(event.TransactionId) },
			"eventId": { S: aws.String(event.Event.EventId) },
			"event": { M: eventItem },
		},
	}

	// Put Item
	_, err = dynamoClient.PutItem(input)

	return err
}
