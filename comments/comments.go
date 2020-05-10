package comments

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/auth"
	"github.com/elliotforbes/api.tutorialedge.net/email"
	"github.com/jinzhu/gorm"
)

// Response a response object
// used for returning an array of comments
type Response struct {
	Comments []Comment `json:"comments"`
}

// Comment the structure
// of comments in the database
type Comment struct {
	gorm.Model
	Slug       string `json:"slug"`
	Body       string `json:"body"`
	Author     string `json:"author"`
	AuthorID   string `json:"sub"`
	Picture    string `json:"picture"`
	ThumbsUp   int    `json:"thumbs_up"`
	ThumbsDown int    `json:"thumbs_down"`
	Heart      int    `json:"heart"`
	Smile      int    `json:"smile"`
}

// GetComments -
// Returns the comments for the given post
func GetComments(request events.APIGatewayProxyRequest, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	fmt.Println(request.QueryStringParameters["slug"])
	slug := request.QueryStringParameters["slug"]

	var comments []Comment
	if slug == "" {
		db.Find(&comments)
	} else {
		db.Where("slug = ?", slug).Find(&comments)
	}

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
func UpdateComment(request events.APIGatewayProxyRequest, tokenInfo auth.TokenInfo, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
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

	if err = db.Save(comment).Error; err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Could not update comment",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	fmt.Println("Comment Saved with new Vote")
	return events.APIGatewayProxyResponse{
		Body:       "Successfull Vote!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// PostComment -
// Adds a new comment to the site
func PostComment(request events.APIGatewayProxyRequest, tokenInfo auth.TokenInfo, db *gorm.DB) (events.APIGatewayProxyResponse, error) {

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

	if err = db.Create(&comment).Error; err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Could not post comment",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

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
func DeleteComment(request events.APIGatewayProxyRequest, tokenInfo auth.TokenInfo, db *gorm.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Deleting a comment")

	fmt.Println(request.QueryStringParameters["id"])
	id := request.QueryStringParameters["id"]

	var comment Comment
	if err := db.First(&comment, id).Error; err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Could not delete comment!",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	if tokenInfo.Sub != comment.AuthorID {
		return events.APIGatewayProxyResponse{
			Body:       "Not Authorized to delete other comment!",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	if err := db.Delete(&comment).Error; err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Could not delete found comment!",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
