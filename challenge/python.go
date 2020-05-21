package challenge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
)

// ExecutePythonChallenge does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecutePythonChallenge(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	var challenge ChallengeRequest
	err := json.Unmarshal(body, &challenge)
	if err != nil {
		fmt.Println("Could not unmarshal challenge")
	}

	err = setupGoEnvironment()
	if err != nil {
		log.Fatalf("Setting up Go Env Failed")
		return events.APIGatewayProxyResponse{
			Body:       "Failed to setup Go Environment",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	out, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())

		log.Fatalf("executing go version failed %s\n", err)
		return events.APIGatewayProxyResponse{
			Body:       "Failed to run go version",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}
	fmt.Printf("Go Version Output: %s", string(out))

	tmpfile, err := ioutil.TempFile("/tmp", "main.*.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created File: " + tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(request.Body)); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	out, err = exec.Command("go", "run", tmpfile.Name()).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       string(out),
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 200,
		}, nil
	}

	fmt.Printf("go run output: %s\n", string(out))

	return events.APIGatewayProxyResponse{
		Body:       string(out),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
