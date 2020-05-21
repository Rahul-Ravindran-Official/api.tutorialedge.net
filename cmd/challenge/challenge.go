package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/challenge"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			Body:       "Not Found",
			StatusCode: 404,
		}, nil
	}

	response, _ := challenge.ExecuteGoChallenge(request)
	return response, nil

	// switch request.Path {
	// case "/api/v1/challengego", "/v1/executego":
	// 	response, _ := challenge.ExecuteGoChallenge(request)
	// 	return response, nil
	// case "/api/v1/executepython", "/v1/executepython":
	// 	response, _ := challenge.ExecutePythonChallenge(request)
	// 	return response, nil
	// default:
	// 	return events.APIGatewayProxyResponse{
	// 		Body:       "Invalid HTTP Method",
	// 		StatusCode: 501,
	// 	}, nil
	// }
}

func main() {
	lambda.Start(handler)
}
