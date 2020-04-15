package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	return events.APIGatewayProxyResponse{
		Body:       "{ \"status\": \"up\" }",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
