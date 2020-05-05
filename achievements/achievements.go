package achievements

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/gorm"
)

type Achievement struct {
	gorm.Model
	Title string `json:"title"`
}

// GetAchievements -
// Returns the Achievements for the given post
func GetAchievements(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Achievements")

	return events.APIGatewayProxyResponse{
		Body:       "Get Request!",
		StatusCode: 200,
	}, nil
}

// UpdateAchievement -
// Updates the Achievement
func UpdateAchievement(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Achievements")

	return events.APIGatewayProxyResponse{
		Body:       "Put Request!",
		StatusCode: 200,
	}, nil
}

// PostAchievement -
// Adds a new Achievement to the site
func PostAchievement(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Achievements")

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		StatusCode: 200,
	}, nil
}

// DeleteAchievement -
// Deletes the Achievement with the ID
func DeleteAchievement(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Achievements")

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		StatusCode: 200,
	}, nil
}

// AllAchievements -
// Returns all Achievements that have been posted to the site
func AllAchievements(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Achievements")

	return events.APIGatewayProxyResponse{
		Body:       "Get All Achievements!",
		StatusCode: 200,
	}, nil
}
