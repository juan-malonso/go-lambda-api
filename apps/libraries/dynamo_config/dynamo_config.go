package dynamo_config

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
var dynamoClient = dynamodb.New(sess)
var dynamoConfigTable = env.GetVar("AWS_DYNAMO_CONFIG_TABLE", "")

// Config Structure
type ConfigValue struct {
	Value string `json:"value"`
}

type Config struct {
	Config ConfigValue `json:"config"`
}

func GetItem(configId string) (*Config, error) {

	// Set Aws Command
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dynamoConfigTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": { S: aws.String(configId) },
		},
	}

	// Get Item
	result, err := dynamoClient.GetItem(input)
	if err != nil { return nil, err }

	item := &Config{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	return item, err
}
