package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type JsonParser struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	Anser string `json:"Anser:"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s := []byte(request.Body)
// 	s := []byte(`{
//     "text": "Jim",
//     "title": "33",
//     "Anser:": "asnnaf"
// }`)
	var data JsonParser
	json.Unmarshal(s, &data)
	fmt.Println("data", data.Text)
	fmt.Println("body", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
