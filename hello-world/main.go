package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	time.Sleep(time.Second * 3)

	uri := url.URL{
		Scheme:      "http",
		Host:        "127.0.0.1:1338",
		Path:        request.Path,
		ForceQuery:  false,
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println(resp.Body)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func runMetadataMock() {
	cmd := exec.Command("/opt/ec2-metadata-mock", "-c", "/opt/config.json")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runMetadataMock()
	lambda.Start(handler)
}

