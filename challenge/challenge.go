package challenge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/TutorialEdge/api.tutorialedge.net/auth"
	"github.com/TutorialEdge/api.tutorialedge.net/email"
	"github.com/TutorialEdge/api.tutorialedge.net/users"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/gorm"
)

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

	var user users.User
	db.Where(users.User{Sub: tokenInfo.Sub}).FirstOrCreate(&user)

	var challenge users.Challenge
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

	err = email.SendEmail("A User Has Completed A Challenge!", challenge.Slug, "admin@tutorialedge.net")
	if err != nil {
		fmt.Println("Error Sending Comment Notification Email...")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
