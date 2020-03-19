package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// GetComments -
// Returns the comments for the given post
func GetComments(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Get Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// UpdateComment -
// Updates the comment
func UpdateComment(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Put Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// PostComment -
// Adds a new comment to the site
func PostComment(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// DeleteComment -
// Deletes the comment with the ID
func DeleteComment(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// AllComments -
// Returns all comments that have been posted to the site
func AllComments(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Get All Comments!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
