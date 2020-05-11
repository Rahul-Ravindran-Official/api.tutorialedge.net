package code_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/elliotforbes/api.tutorialedge.net/code"
)

func TestExecuteCode(t *testing.T) {
	body := []byte(`package main

	import "fmt"

	func main() {

		mymap := make(map[string]int)

		mymap["elliot"] = 25

		// we can use this if statement to check to see if
		// a given key "elliot" exists within a map in Go
		if _, ok := mymap["elliot"]; ok {
			// the key 'elliot' exists within the map
			fmt.Println(mymap["elliot"])
		}
	}
	`)

	request := events.APIGatewayProxyRequest{}
	request.Body = string(body)

	response, err := code.ExecuteCode(request)
	if err != nil {
		t.Fail()
	}
	fmt.Println(response)

}
