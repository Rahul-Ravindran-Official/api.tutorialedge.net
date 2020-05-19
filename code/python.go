package code

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
)


// ExecutePython takes in a request, retreives the python code
// from the body of that request and then executes it.
func ExecutePython(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	tmpfile, err := ioutil.TempFile("/tmp", "main.*.py")
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


	out, err := exec.Command("python", tmpfile.Name()).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       string(out),
			Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
			StatusCode: 200,
		}, nil
	}

	fmt.Printf("python output: %s\n", string(out))

	return events.APIGatewayProxyResponse{
		Body:       string(out),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil

}
