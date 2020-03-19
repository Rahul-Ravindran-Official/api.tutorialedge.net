package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	if request.HTTPMethod == "GET" {
		response, _ := GetPosts(request)
		return response, nil
	} else if request.HTTPMethod == "POST" {
		response, _ := PostPost(request)
		return response, nil
	} else if request.HTTPMethod == "PUT" {
		response, _ := UpdatePost(request)
		return response, nil
	} else if request.HTTPMethod == "DELETE" {
		response, _ := DeletePost(request)
		return response, nil
	} else {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid HTTP Method",
			StatusCode: 501,
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}