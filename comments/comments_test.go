// +build integration

package comments

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/comments"
	"github.com/jinzhu/gorm"
)

func getDBConn() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := 25060
	dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable

	db, err := gorm.Open("mysql", dbConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func TestPostComments(t *testing.T) {
	fmt.Println("Testing Post Comments...")
	// t.Error()
}

func TestDeleteComments(t *testing.T) {
	fmt.Println("Testing Delete Comments...")
	// t.Error()
}

func TestRetrieveComments(t *testing.T) {
	fmt.Println("Testing We can retrieve all comments")

	db := getDBConn()

	response := comments.AllComments(&events.APIGatewayProxyRequest{}, db)

	if response.StatusCode != 200 {
		fmt.Println("Failed to retrieve all comments...")
		t.Error()
	}
}
