package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/urandom/text-summary/summarize"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

type briefStruct struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	Brief string `json:"brief"`
}

func responseJson(title string, text string, brief string) string {
	res := briefStruct{
		Title: title,
		Text:  text,
		Brief: brief,
	}

	marshaled, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return "Marshall error"
	}

	return string(marshaled)
}

func briefText(title string, text string) string {
	s := summarize.NewFromString(title, text)

	return strings.Join(s.KeyPoints(), " ")
}

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data briefStruct
	json.Unmarshal([]byte(request.Body), &data)

	res := responseJson(data.Title, data.Text, briefText(data.Title, data.Text))
	fmt.Println("res: ", res)
	return events.APIGatewayProxyResponse{Body: res, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
