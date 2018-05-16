package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

// ErrNameNotProvided represents a missing name error
var ErrNameNotProvided = errors.New("name missing from request payload")

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer instrument("main.handler")()

	if len(req.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + req.Body,
		StatusCode: http.StatusOK,
	}
}

func instrument(tag string) func() {
	callTime := time.Now()
	return func() {
		log.WithFields(log.Fields{
			"method":         tag,
			"execution_time": time.Since(callTime),
		}).Infof("%s took %d to execute", tag, time.Since(callTime))
	}
}

func main() {
	lambda.Start(Handler)
}
