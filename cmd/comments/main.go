package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/auth"
	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/elliotforbes/api.tutorialedge.net/database"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	db, err := database.GetDBConn()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	switch request.HTTPMethod {
	case "GET":
		response, _ := comments.GetComments(request, db)
		return response, nil
	case "POST":
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.PostComment(request, db)
			return response, nil
		}
		return events.APIGatewayProxyResponse{
			Body:       "Not Authorized",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	case "PUT":
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.UpdateComment(request, db)
			return response, nil
		}

		return events.APIGatewayProxyResponse{
			Body:       "Not Authorized",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	case "DELETE":
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.DeleteComment(request, db)
			return response, nil
		}

		return events.APIGatewayProxyResponse{
			Body:       "Not Authorized",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil

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
