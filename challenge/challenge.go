package challenge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/TutorialEdge/api.tutorialedge.net/auth"
	"github.com/TutorialEdge/api.tutorialedge.net/email"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/gorm"
)

// Challenge - holds the users challenges
type Challenge struct {
	gorm.Model
	Slug          string `json:"slug"`
	AuthorID      string `json:"sub"`
	Code          string `json:"code"`
	Score         int    `json:"score"`
	Passed        bool   `json:"passed"`
	ExecutionTime string `json:"execution_time"`
}

// Response a response object
// used for returning an array of challenges
type Response struct {
	Challenges []Challenge `json:"challenges"`
}

// GetChallenge - Retrieves a challenge
func GetChallenge(request events.APIGatewayProxyRequest, tokenInfo auth.TokenInfo, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Retrieving Challenge State for User")

	fmt.Println(request.QueryStringParameters["slug"])
	slug := request.QueryStringParameters["slug"]

	var challenges []Challenge
	if slug == "" {
		db.Find(&challenges)
	} else {
		db.Where("slug = ?", slug).Find(&challenges)
	}

	response := Response{
		challenges: challenges,
	}

	fmt.Printf("%+v\n", challenges)

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

	var challenge Challenge
	err = json.Unmarshal([]byte(request.Body), &challenge)
	if err != nil {
		panic(err.Error())
	}

	if db.Where("author_id = ? AND slug = ?", tokenInfo.Sub, challenge.Slug).Find(&challenge).RecordNotFound() {
		fmt.Println("Challenge not already completed, progressing")

		challenge.AuthorID = tokenInfo.Sub

		if err = db.Create(&challenge).Error; err != nil {
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
			Body:       "Challenge Saved!",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 200,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			Body:       "Challenge Already Saved!",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 200,
		}, nil
	}

}
