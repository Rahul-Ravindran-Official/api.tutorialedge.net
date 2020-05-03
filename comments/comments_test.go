// +build integration

package comments_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/elliotforbes/api.tutorialedge.net/database"
)

func TestGetComments(t *testing.T) {
	db, err := database.GetDBConn()
	if err != nil {
		t.Log(err)
		t.Error("Could not get DB Connection")
		return
	}

	request := events.APIGatewayProxyRequest{}

	request.QueryStringParameters["slug"] = "/test/"

	response, err := comments.AllComments(request, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("Failed to retrieve all comments...")
		t.Error("Retrieving all comments returned unexpected status code")
	}
}

func TestUpdateComments(t *testing.T) {
	fmt.Println("Testing Updates to comments")
	// t.Error()
}

func TestPostComments(t *testing.T) {
	fmt.Println("Testing Post Comments...")
	// t.Error()
}

func TestDeleteComments(t *testing.T) {
	fmt.Println("Testing Delete Comments...")
	// t.Error()
}

func TestRetrieveComments(t *testing.T) {
	fmt.Println("Testing We can retrieve all comments")

	db, err := database.GetDBConn()
	if err != nil {
		t.Log(err)
		t.Error("Could not get DB Connection")
		return
	}

	response, err := comments.GetComments(events.APIGatewayProxyRequest{}, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("Failed to retrieve all comments...")
		t.Error("Retrieving all comments returned unexpected status code")
	}
}
