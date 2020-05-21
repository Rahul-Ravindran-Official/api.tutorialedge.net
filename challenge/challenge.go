package challenge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mholt/archiver/v3"
)

// CodeResponse contains the response from
// executing the Go code
type CodeResponse struct {
	ExitCode string `json:"exit_code"`
	Output   string `json:"output"`
}

// ChallengeRequest takes in the source code
// from the editor as well as a number of tests
// which are written into a file and ran
type ChallengeRequest struct {
	Code  string          `json:"code"`
	Tests []ChallengeTest `json:"tests"`
}

// ChallengeTest is a struct which contains
// the source code for a test file as well as
// the metadata such as the test name
type ChallengeTest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func setupGoEnvironment() error {
	path := os.Getenv("PATH")
	os.Setenv("PATH", path+":/tmp/go/bin")
	os.Setenv("GOROOT", "/tmp/go")
	os.Setenv("GOPATH", "/tmp")
	os.Setenv("GOCACHE", "/tmp/go-cache")

	if _, err := os.Stat("/tmp/go"); os.IsNotExist(err) {
		err := os.Mkdir("/tmp/go", 0777)
		if err != nil {
			fmt.Println(err)
		}

		// untar ./code/go.tar.gz -> /tmp/go
		err = archiver.Unarchive("./code/go.tar.gz", "/tmp")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

// ExecuteGoChallenge does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecuteGoChallenge(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var challenge ChallengeRequest
	err := json.Unmarshal(request.Body, &challenge)
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

	if _, err := tmpfile.Write([]byte(challenge.Code)); err != nil {
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
