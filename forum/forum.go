package forum

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// GetPosts -
// Returns the Posts for the given post
func GetPosts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Posts")

	return events.APIGatewayProxyResponse{
		Body:       "Get Request!",
		StatusCode: 200,
	}, nil
}

// UpdatePost -
// Updates the Post
func UpdatePost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Posts")

	return events.APIGatewayProxyResponse{
		Body:       "Put Request!",
		StatusCode: 200,
	}, nil
}

// PostPost -
// Adds a new Post to the site
func PostPost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Posts")

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		StatusCode: 200,
	}, nil
}

// DeletePost -
// Deletes the Post with the ID
func DeletePost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Posts")

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		StatusCode: 200,
	}, nil
}

// AllPosts -
// Returns all Posts that have been posted to the site
func AllPosts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Posts")

	return events.APIGatewayProxyResponse{
		Body:       "Get All Posts!",
		StatusCode: 200,
	}, nil
}
