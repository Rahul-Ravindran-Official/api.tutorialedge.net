package comments

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/email"
	"github.com/jinzhu/gorm"
)

type BodyRequest struct {
	RequestName string `json:"name"`
}

type Response struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	Posted      string `json:"posted"`
	Picture     string `json:"picture,omitempty"`
	Thumbs_up   int    `json:"thumbs_up,omitempty"`
	Thumbs_down int    `json:"thumbs_down,omitempty"`
	Heart       int    `json:"heart,omitempty"`
	Smile       int    `json:"smile,omitempty"`
}

type Vote struct {
	Id   int    `json:"id"`
	Vote string `json:"vote"`
}

// GetComments -
// Returns the comments for the given post
func GetComments(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	fmt.Println(request.QueryStringParameters["slug"])
	slug := request.QueryStringParameters["slug"]

	var comments []Comment
	db.Where("slug = ?", slug).Find(&comments)

	response := Response{
		Comments: comments,
	}

	fmt.Printf("%+v\n", comments)

	jsonResults, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", string(jsonResults))

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResults),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// UpdateComment -
// Updates the comment
func UpdateComment(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Adding a Vote to a Comment")
	fmt.Printf("Request: %v\n", request)

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	var comment Comment

	err := json.Unmarshal(body, &comment)
	if err != nil {
		panic(err.Error())
	}

	db.Save(comment)
	fmt.Println("Comment Saved with new Vote")

	return events.APIGatewayProxyResponse{
		Body:       "Successfull Vote!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// PostComment -
// Adds a new comment to the site
func PostComment(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Posting a Comment")
	fmt.Println("Received body: ", request.Body)
	body, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))

	var comment Comment

	err = json.Unmarshal([]byte(request.Body), &comment)
	if err != nil {
		panic(err.Error())
	}

	db.Create(comment)

	err = email.SendEmail("A New Comment Has been Posted!", comment.Body, "admin@tutorialedge.net")
	if err != nil {
		fmt.Println("Error Sending Comment Notification Email...")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// DeleteComment -
// Deletes the comment with the ID
func DeleteComment(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// AllComments -
// Returns all comments that have been posted to the site
func AllComments(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	var comments []Comment

	db.Find(&comments)

	jsonResults, err := json.Marshal(comments)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", string(jsonResults))

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResults),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
