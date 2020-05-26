package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tutorialedge/api.tutorialedge.net/auth"
	"github.com/tutorialedge/api.tutorialedge.net/comments"
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
	case "GET":
		response, _ := comments.GetComments(request, db)
		return response, nil
	case "POST":
		if ok, tokenInfo := auth.Authenticate(request); ok {
			response, _ := comments.PostComment(request, tokenInfo, db)
			return response, nil
		}
		return auth.UnauthorizedResponse(), nil
	case "PUT":
		if ok, tokenInfo := auth.Authenticate(request); ok {
			response, _ := comments.UpdateComment(request, tokenInfo, db)
			return response, nil
		}
		return auth.UnauthorizedResponse(), nil
	case "DELETE":
		if ok, tokenInfo := auth.Authenticate(request); ok {
			response, _ := comments.DeleteComment(request, tokenInfo, db)
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
