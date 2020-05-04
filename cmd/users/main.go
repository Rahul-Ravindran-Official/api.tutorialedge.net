package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/database"
	"github.com/elliotforbes/api.tutorialedge.net/users"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := database.GetDBConn()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	switch request.HTTPMethod {
	case "GET":
		response, _ := users.GetUser(request, db)
		return response, nil
	case "POST":
		response, _ := users.NewUser(request, db)
		return response, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "{\"status\": \"success\"}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
