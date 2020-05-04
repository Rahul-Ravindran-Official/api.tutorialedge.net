// +build integration

package users_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/elliotforbes/api.tutorialedge.net/database"
	"github.com/elliotforbes/api.tutorialedge.net/users"
)

func TestGetUser(t *testing.T) {
	db, err := database.GetDBConn()
	if err != nil {
		t.Log(err)
		t.Error("Could not get DB Connection")
		return
	}

	request := events.APIGatewayProxyRequest{}
	request.QueryStringParameters = make(map[string]string)
	request.QueryStringParameters["name"] = "Elliot Forbes"

	response, err := users.GetUser(request, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("Failed to retrieve all users...")
		t.Error("Retrieving all users returned unexpected status code")
	}
}