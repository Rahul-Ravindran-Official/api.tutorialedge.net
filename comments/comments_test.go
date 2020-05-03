// +build integration

package comments

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/elliotforbes/api.tutorialedge.net/database"
)

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
		t.Error("Could not get DB Connection")
	}

	response := comments.AllComments(&events.APIGatewayProxyRequest{}, db)

	if response.StatusCode != 200 {
		fmt.Println("Failed to retrieve all comments...")
		t.Error("Retrieving all comments returned unexpected status code")
	}
}
