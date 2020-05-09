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
	request.QueryStringParameters = make(map[string]string)

	request.QueryStringParameters["slug"] = "/test/"

	response, err := comments.GetComments(request, db)
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
	db, err := database.GetDBConn()
	if err != nil {
		t.Log(err)
		t.Error("Could not get DB Connection")
		return
	}

	body := []byte(`{
		"slug": "/random-test-slug/",
		"body": "Test",
		"author": "Elliot Forbes",
		"picture": "https://lh3.googleusercontent.com/a-/AOh14GjvaxQsEaZVJi809M65ACo6BogR8nI-Ntl7FuBkEA",
		"user": {
			"sub": "google-oauth2|116080913132292461306",
			"given_name": "Elliot",
			"family_name": "Forbes",
			"nickname": "efgamercity",
			"name": "Elliot Forbes",
			"picture": "https://lh3.googleusercontent.com/a-/AOh14GjvaxQsEaZVJi809M65ACo6BogR8nI-Ntl7FuBkEA",
			"locale": "en-GB",
			"updated_at": "2020-05-04T08:03:18.105Z"
		}
	}`)

	request := events.APIGatewayProxyRequest{}
	request.Body = string(body)

	response, err := comments.PostComment(request, db)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(response)

	var comment comments.Comment
	db.Where("slug = ?", "/random-test-slug/").Find(&comment)

	if comment.Body != "Test" {
		t.Fail()
	}

	// 	db.Delete(&comment)
}

func TestFailedDeleteComments(t *testing.T) {
	fmt.Println("Testing Delete Comments...")
	// t.Error()
	db, err := database.GetDBConn()
	if err != nil {
		t.Log(err)
		t.Error("Could not get DB Connection")
		return
	}

	request := events.APIGatewayProxyRequest{}
	request.QueryStringParameters = make(map[string]string)

	request.QueryStringParameters["id"] = "-1"

	response, err := comments.DeleteComment(request, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response.Body)

	if response.StatusCode != 503 {
		fmt.Println("Did not receive expected 503 response")
		t.Error("Failed Delete Comment test did not return 503 expected")
	}

}

// func TestDeleteComments(t *testing.T) {
// 	fmt.Println("Testing Delete Comments...")
// 	// t.Error()
// 	db, err := database.GetDBConn()
// 	if err != nil {
// 		t.Log(err)
// 		t.Error("Could not get DB Connection")
// 		return
// 	}

// 	request := events.APIGatewayProxyRequest{}
// 	request.QueryStringParameters = make(map[string]string)

// 	request.QueryStringParameters["id"] = "29"

// 	response, err := comments.DeleteComment(request, db)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	t.Log(response.Body)

// 	if response.StatusCode != 200 {
// 		fmt.Println("Failed to delete comment...")
// 		t.Error("Deleting comment returned unexpected status code")
// 	}

// }

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
