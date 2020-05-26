package main

import (
	"fmt"

	"github.com/TutorialEdge/api.tutorialedge.net/auth"
	"github.com/TutorialEdge/api.tutorialedge.net/challenge"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tutorialedge/api.tutorialedge.net/database"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	db, err := database.GetDBConn()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	switch request.HTTPMethod {
	case "POST":
		if ok, tokenInfo := auth.Authenticate(request); ok {
			response, _ := challenge.PostChallenge(request, db)
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
