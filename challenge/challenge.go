package challenge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/gorm"
	"github.com/tutorialedge/api.tutorialedge.net/auth"
	"github.com/tutorialedge/api.tutorialedge.net/email"
	"github.com/tutorialedge/api.tutorialedge.net/user"
)

type Challenge struct {
	gorm.Model
	Slug          string `json:"slug"`
	Code          string `json:"code"`
	Score         int    `json:"score"`
	Passed        bool   `json:"passed"`
	ExecutionTime string `json:"execution_time"`
}

// PostChallenge - Adds a challenge to a User entry in the database
//
func PostChallenge(request events.APIGatewayProxyRequest, tokenInfo auth.TokenInfo, db *gorm.DB) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Posting a Challenge")
	fmt.Println("Received body: ", request.Body)

	body, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))

	if tokenInfo.Sub == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Could not post challenge with no Sub",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	var user user.User
	db.Where(User{Sub: tokenInfo.Sub}).FirstOrCreate(&user)

	var challenge Challenge
	err = json.Unmarshal([]byte(request.Body), &challenge)
	if err != nil {
		panic(err.Error())
	}

	user.Challenges = append(user.Challenges, challenge)

	if er = db.Save(&user).Error; err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Could not save challenge for user",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	err = email.SendEmail("A User Has Completed A Challenge!", comment.Slug, "admin@tutorialedge.net")
	if err != nil {
		fmt.Println("Error Sending Comment Notification Email...")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
