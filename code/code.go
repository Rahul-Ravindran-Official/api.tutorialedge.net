package code

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
)

// CodeResponse contains the response from
// executing the Go code
type CodeResponse struct {
	ExitCode string `json:"exit_code"`
	Output   string `json:"output"`
}

// ExecuteCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecuteCode(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	out, err := exec.Command("./bin/go", "version").Output()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(out)

	// the WriteFile method returns an error if unsuccessful
	// err = ioutil.WriteFile("temp/main.go", body, 0777)
	// // handle this error
	// if err != nil {
	// 	// print it out
	// 	fmt.Println(err)
	// }

	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
