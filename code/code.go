package code

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
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

func ExecuteCode(request events.APIGatewayProxyRequest) {

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("temp/main.go", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	cmd := exec.Command("go", "version")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResults),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
