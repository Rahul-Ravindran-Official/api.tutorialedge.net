package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	buildNumber := os.Getenv("TRAVIS_BUILD_NUMBER")

	return events.APIGatewayProxyResponse{
		Body:       "{ \"status\": \"up\", \"version\": \"" + buildNumber + "\"}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
