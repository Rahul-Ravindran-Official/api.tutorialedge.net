package code

import (
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

	if _, err := os.Stat("/tmp/go"); os.IsNotExist(err) {
		err := os.Mkdir("/tmp/go", 0777)
		if err != nil {
			fmt.Println(err)
		}

		// untar ./code/go.tar.gz -> /tmp/go
		output, err := exec.Command("tar", "-xzf", "./code/go.tar.gz", "-C", "/tmp").CombinedOutput()
		if err != nil {
			fmt.Println("Failed to Execute tar command")
			fmt.Println(err.Error())
			fmt.Println(string(output))
			return err
		}
	}
	return nil
}

// ExecuteCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecuteCode(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)
	// body, err := base64.StdEncoding.DecodeString(request.Body)
	// if err != nil {
	// 	fmt.Println("Issue decoding request body from base64")
	// }
	// fmt.Println(string(body))

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
