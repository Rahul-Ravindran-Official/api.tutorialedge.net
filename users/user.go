package users

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/elliotforbes/api.tutorialedge.net/email"
	"github.com/jinzhu/gorm"
)

// User - A user on TutorialEdge! :D
type User struct {
	gorm.Model
	Name     string             `json:"name"`
	Comments []comments.Comment `json:"comments"`
}

// GetComments - returns all the comments that a
// user has posted to the site
func GetUser(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Retrieving all Comments by User")
	fmt.Println(request.QueryStringParameters["name"])
	name := request.QueryStringParameters["name"]

	var comments []comments.Comment

	db.Where("author = ?", name).Find(&comments)

	var user User
	user.Name = name
	user.Comments = comments

	jsonResults, err := json.Marshal(user)
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

// NewUser - handles what actions should be taken if a new user
// is registered on the site
func NewUser(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	err := email.SendEmail("New User Account Registered!", "A New User has registered on TutorialEdge", "admin@tutorialedge.net")

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "{\"status\": \"error: " + err.Error() + "\"}",
			StatusCode: 503,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "{\"status\": \"success\"}",
		StatusCode: 200,
	}, nil
}
