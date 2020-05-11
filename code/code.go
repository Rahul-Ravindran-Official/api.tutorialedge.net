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

// ExecuteCode does the job of taking the Go code that has
// been sent to API from a snippet and executing it before
// returning the response
func ExecuteCode(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	path := os.Getenv("PATH")
	os.Setenv("PATH", path+":"+os.Getenv("LAMBDA_TASK_ROOT")+"/bin")

	out, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())

		out, _ := exec.Command("pwd").Output()
		fmt.Println(string(out))

		out, _ = exec.Command("env").Output()
		fmt.Println(string(out))

		out, _ = exec.Command("where", "go").Output()
		fmt.Println(string(out))

		out, _ = exec.Command("type", "go").Output()
		fmt.Println(string(out))

		out, _ = exec.Command("ls", "-ltr", "bin").Output()
		fmt.Println(string(out))

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
