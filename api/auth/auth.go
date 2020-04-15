package auth

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Authenticate(request events.APIGatewayProxyRequest) {
	fmt.Println("Attempting to Authenticate Incoming Request...")

	token := request.Headers["Authorization"]

	fmt.Println(token)
}
