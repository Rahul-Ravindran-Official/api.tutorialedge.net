package auth

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	fmt.Println("Testing Authenticate Function")
}

func TestBadToken(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  bool
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{},
			expect:  false,
		},
	}

	for _, test := range tests {
		response := Authenticate(test.request)
		assert.Equal(t, test.expect, response)
	}

}
