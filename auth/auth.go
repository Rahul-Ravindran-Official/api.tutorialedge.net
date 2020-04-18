package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

func Authenticate(request events.APIGatewayProxyRequest) bool {
	fmt.Println("Attempting to Authenticate Incoming Request...")

	header := request.Headers["Authorization"]
	tokenString := strings.Split(string(header), " ")[1]

	signingKey := os.Getenv("AUTH0_SIGNING_KEY")

	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(signingKey))
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