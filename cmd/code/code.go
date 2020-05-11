package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/auth"
	"github.com/elliotforbes/api.tutorialedge.net/code"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	switch request.HTTPMethod {
	case "POST":
		if ok, tokenInfo := auth.Authenticate(request); ok {
			response, _ := code.ExecuteCode(request)
			return response, nil
		}
		return auth.UnauthorizedResponse(), nil
	default:
		return events.APIGatewayProxyResponse{
			Body:       "Invalid HTTP Method",
			StatusCode: 501,
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
