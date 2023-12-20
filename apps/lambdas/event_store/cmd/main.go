package main

import (
	"context"
  "encoding/json"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/go-lambda-api/apps/libraries/commons"
	"github.com/go-lambda-api/apps/libraries/log"
	"github.com/go-lambda-api/apps/libraries/dynamo_events"
)

func DynamoEvents(event interface{}) (dynamo_events.Event, error) {
	var dynamoEvent dynamo_events.Event

	jsonEvent, err := json.Marshal(event)
	if err != nil { return dynamoEvent, err }

	err = json.Unmarshal(jsonEvent, &dynamoEvent)
	return dynamoEvent, err
}

var logger = log.InitLogger("endpoint_config")

func HandleRequest(ctx context.Context, req commons.Event) (commons.Event, error) {

	// =================================================================== LOG ==

	logger.WithFields(log.Fields{ "req": req }).Info("STORE_EVENT_REQUEST")

	// ==================================================== DYNAMO STORE EVENT ==

	event, err := DynamoEvents(req.Body)
	if err != nil { return commons.Event{}, err }

	logger.WithFields(log.Fields{ "event": event }).Info()

	err = dynamo_events.PutItem(event);
	if err != nil { return commons.Event{}, err }

	// ============================================================== RESPONSE ==

	return commons.Event{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
