package comments

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/elliotforbes/api.tutorialedge.net/auth"
	"github.com/elliotforbes/api.tutorialedge.net/comments"
	_ "github.com/go-sql-driver/mysql"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := 25060
	dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable

	db, err := sql.Open("mysql", dbConnectionString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if request.HTTPMethod == "GET" {
		response, _ := comments.GetComments(request, db)
		return response, nil
	} else if request.HTTPMethod == "POST" {
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.PostComment(request, db)
			return response, nil
		} else {
			return events.APIGatewayProxyResponse{
				Body:       "Not Authorized",
				Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
				StatusCode: 503,
			}, nil
		}
	} else if request.HTTPMethod == "PUT" {
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.UpdateComment(request, db)
			return response, nil
		} else {
			return events.APIGatewayProxyResponse{
				Body:       "Not Authorized",
				Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
				StatusCode: 503,
			}, nil
		}
	} else if request.HTTPMethod == "DELETE" {
		if ok := auth.Authenticate(request); ok {
			response, _ := comments.DeleteComment(request, db)
			return response, nil
		} else {
			return events.APIGatewayProxyResponse{
				Body:       "Not Authorized",
				Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
				StatusCode: 503,
			}, nil
		}
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
