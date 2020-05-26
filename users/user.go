package users

import (
	"encoding/json"
	"fmt"

	"github.com/TutorialEdge/api.tutorialedge.net/comments"
	"github.com/TutorialEdge/api.tutorialedge.net/email"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/gorm"
)

// User - A user on TutorialEdge! :D
type User struct {
	gorm.Model
	Name       string             `json:"name"`
	Sub        string             `json:"sub"`
	GivenName  string             `json:"given_name"`
	FamilyName string             `json:"family_name"`
	Nickname   string             `json:"nickname"`
	Picture    string             `json:"picture"`
	Aud        string             `json:"aud"`
	Locale     string             `json:"locale"`
	UpdatedAt  string             `json:"update_at"`
	Comments   []comments.Comment `json:"comments"`
	Challenges []Challenge        `json:"challenges"`
}

// Challenge - holds the users challenges
type Challenge struct {
	Slug          string `json:"slug"`
	Code          string `json:"code"`
	Score         int    `json:"score"`
	Passed        bool   `json:"passed"`
	ExecutionTime string `json:"execution_time"`
}

// Result
type Result struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

// GetComments - returns all the comments that a
// user has posted to the site
func GetUser(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Retrieving all Comments by User")
	fmt.Println(request.QueryStringParameters["sub"])
	sub := request.QueryStringParameters["sub"]

	var comments []comments.Comment
	db.Where("author_id = ?", sub).Find(&comments)
	fmt.Printf("%+v\n", comments)

	var challenges []Challenge
	db.Where("author_id = ?", sub).Find(&challenges)
	fmt.Printf("%+v\n", challenges)

	var user User
	user.Sub = sub
	user.Comments = comments
	user.Challenges = challenges

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
