package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/email"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	email.SendNewUserEmail("New User Account Registered!", "A New User has registered on TutorialEdge", "admin@tutorialedge.net")

	return events.APIGatewayProxyResponse{
		Body:       "{\"status\": \"success\"}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
