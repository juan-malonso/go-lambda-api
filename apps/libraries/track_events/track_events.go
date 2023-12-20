package track_events

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"

	"github.com/go-lambda-api/apps/libraries/env"
	"github.com/go-lambda-api/apps/libraries/dynamo_events"
)

// Create Aws Session
var sess, sessErr = session.NewSession(&aws.Config{
	Endpoint:    aws.String(env.GetVar("AWS_DYNAMO_ENDPOINT", "")),
})

// Create Dynamo Client
var eventbridgeClient = eventbridge.New(sess)
var eventbridgeEventsBusName = env.GetVar("AWS_EVENTBRIDGE_EVENTS_BUS_NAME", "")

// Event Structure

type Event = dynamo_events.Event

func Send(event Event) error {
	detail, err := dynamo_events.Marshal(event)

	// Send Message
	_, err = eventbridgeClient.PutEvents(&eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				Detail:     aws.String(fmt.Sprintf("%v", detail)),
				DetailType: aws.String("MyAppEvent"),
				Source:     aws.String("com.mycompany.myapp"),
				EventBusName: aws.String(eventbridgeEventsBusName), // Replace with your event bus name if you have a custom one
			},
		},
	})

	return err
}
