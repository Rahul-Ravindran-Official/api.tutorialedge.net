package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	buildNumber := os.Getenv("TRAVIS_BUILD_NUMBER")
	return events.APIGatewayProxyResponse{
		Body:       "{ \"status\": \"up\", \"version\": \"" + buildNumber + "\"}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
