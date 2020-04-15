package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

var mySigningKey = []byte(``)

func Authenticate(request events.APIGatewayProxyRequest) bool {
	fmt.Println("Attempting to Authenticate Incoming Request...")

	header := request.Headers["Authorization"]
	tokenString := strings.Split(string(header), " ")[1]

	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(mySigningKey)
		if err != nil {
			panic(err.Error())
		}
		return verifyKey, nil
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(token)

	if token.Valid {
		return true
	}

	return false
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("%+v\n", request)

	var dbConnectionString string
	dbConnectionString = os.Getenv("DB_CONN_STRING")

	db, err := sql.Open("mysql", dbConnectionString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if request.HTTPMethod == "GET" {
		response, _ := GetComments(request, db)
		return response, nil
	} else if request.HTTPMethod == "POST" {
		if ok := Authenticate(request); ok {
			response, _ := PostComment(request, db)
			return response, nil
		} else {
			return events.APIGatewayProxyResponse{
				Body:       "Not Authorized",
				Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
				StatusCode: 503,
			}, nil
		}
	} else if request.HTTPMethod == "PUT" {
		if ok := Authenticate(request); ok {
			response, _ := UpdateComment(request, db)
			return response, nil
		} else {
			return events.APIGatewayProxyResponse{
				Body:       "Not Authorized",
				Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
				StatusCode: 503,
			}, nil
		}
	} else if request.HTTPMethod == "DELETE" {
		if ok := Authenticate(request); ok {
			response, _ := DeleteComment(request, db)
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
