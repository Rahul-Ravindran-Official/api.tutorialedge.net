package challenge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

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
	Name   string `json:"name"`
	Code   string `json:"code"`
	Test   string `json:"test"`
	Output string `json:"output"`
	Passed bool   `json:"passed"`
}

// ChallengeResponse is the struct that contains the
// response sent back when a challenge is attempted
type ChallengeResponse struct {
	Tests  []ChallengeTest `json:"tests"`
	Built  bool            `json:"built"`
	Output string          `json:"output"`
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
	err := json.Unmarshal([]byte(request.Body), &challenge)
	if err != nil {
		fmt.Println("Could not unmarshal challenge")
		fmt.Println(err.Error())
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

	dir, err := ioutil.TempDir("/tmp", "challenge*")
	if err != nil {
		log.Fatal(err)
	}

	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(challenge.Code), 0666); err != nil {
		log.Fatal(err)
	}
	var response ChallengeResponse

	out, err = exec.Command("go", "run", tmpfn).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		response.Output = string(out)
		response.Built = false
		jsonResults, err := json.Marshal(response)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%+v\n", string(jsonResults))

		return events.APIGatewayProxyResponse{
			Body:       string(jsonResults),
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 200,
		}, nil
	}

	response.Output = string(out)
	response.Built = true

	for _, test := range challenge.Tests {
		tmpfn := filepath.Join(dir, test.Name+".go")
		if err := ioutil.WriteFile(tmpfn, []byte(test.Code), 0666); err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("go", "test", "-run", test.Test)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			test.Output = err.Error()
			test.Passed = false
		} else {
			test.Output = string(out)
			test.Passed = true
		}

		response.Tests = append(response.Tests, test)
		fmt.Printf("go test %s\n", tmpfn)
		fmt.Printf("%+v\n", string(out))
	}

	fmt.Printf("go run output: %s\n", string(out))
	jsonResults, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", string(jsonResults))

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResults),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
