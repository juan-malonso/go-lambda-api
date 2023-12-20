package main

import (
	"context"
	"encoding/json"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/go-lambda-api/apps/libraries/commons"
	"github.com/go-lambda-api/apps/libraries/log"
	"github.com/go-lambda-api/apps/libraries/dynamo_config"
)

var subdomain = regexp.MustCompile(`^http[s]?:\/\/([\w-]+)[.]?.*(:\d*)?`)
var logger = log.InitLogger("endpoint_config")

type Request struct {
	Body interface{} `json:"body,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Query interface{} `json:"querystring,omitempty"`
	Path interface{} `json:"path,omitempty"`
}


func HandleRequestHeaders(raw interface{}) (commons.Headers, error) {
	var headers commons.Headers

	jsonHeaders, err := json.Marshal(raw)
	if err != nil { return headers, err }

	err = json.Unmarshal(jsonHeaders, &headers)
	return headers, err
}

func HandleRequestBody(raw interface{}) (commons.Body, error) {
	var body commons.Body

	jsonBody, err := json.Marshal(raw)
	if err != nil { return body, err }

	err = json.Unmarshal(jsonBody, &body)
	return body, err
}

func HandleRequest(ctx context.Context, req Request) (commons.Event, error) {

	logger.WithFields(log.Fields{ "req": req }).Info("REQUEST")

	headers, err := HandleRequestHeaders(req.Headers)
	if err != nil { return commons.Event{}, err }

	body, err := HandleRequestBody(req.Body)
	if err != nil { return commons.Event{}, err }

	// ============================================================ LANDING ID ==

	configId := ""
	match := subdomain.FindStringSubmatch(headers.Origin)
	if len(match) >= 2 { configId = match[1] }

	// ========================================================= DYNAMO CONFIF ==

	logger.WithFields(log.Fields{ "configId": configId }).Info("GET_LANDING_CONFIG_REQUEST")
	config, err := dynamo_config.GetItem(configId);
	if err != nil { return commons.Event{}, err }

	// ============================================================== RESPONSE ==

	return commons.Event{
		StatusCode: 200,

		ConfigId: configId,
		Config: config,
		
		Body: body,
		Headers: headers,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
