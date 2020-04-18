package achievements

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/elliotforbes/api.tutorialedge.net/achievements"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	if request.HTTPMethod == "GET" {
		response, _ := achievements.GetAchievements(request)
		return response, nil
	} else if request.HTTPMethod == "POST" {
		response, _ := achievements.PostAchievement(request)
		return response, nil
	} else if request.HTTPMethod == "PUT" {
		response, _ := achievements.UpdateAchievement(request)
		return response, nil
	} else if request.HTTPMethod == "DELETE" {
		response, _ := achievements.DeleteAchievement(request)
		return response, nil
	} else {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid HTTP Method",
			StatusCode: 501,
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
