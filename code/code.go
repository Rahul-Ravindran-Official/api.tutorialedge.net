package code

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
)

// CodeResponse contains the response from
// executing the Go code
type CodeResponse struct {
	ExitCode string `json:"exit_code"`
	Output   string `json:"output"`
}

func setupGoEnvironment() error {
	path := os.Getenv("PATH")
	os.Setenv("PATH", path+":/tmp/go/bin")
	os.Setenv("GOROOT", "/tmp/go")
	os.Setenv("GOPATH", "/tmp")
	os.Setenv("GOCACHE", "/tmp/go-cache")

	r, err := os.Open("./code/go.tar.gz")
	if err != nil {
		fmt.Println("error")
		return err
	}
	// untar ./code/go.tar.gz -> /tmp/go
	out, err := exec.Command("tar", "-C", "/tmp/go", "-xzf", "./code/go.tar.gz").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// ExecuteCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecuteCode(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	err := setupGoEnvironment()
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
	fmt.Println(string(out))

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("main.go", body, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       "Failed to write main.go",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	out, err = exec.Command("go", "run", "main.go").CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       "Failed to run main.go",
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 503,
		}, nil
	}

	fmt.Println(string(out))

	return events.APIGatewayProxyResponse{
		Body:       string(out),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
