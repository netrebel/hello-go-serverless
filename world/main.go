package main

import (
	"bytes"
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

type MyEvent struct {
	Body string `json:"body"`
}

type Body struct {
	Name string `json:"name"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(event MyEvent) (Response, error) {
	var buf bytes.Buffer

	fmt.Printf("event: %+v\n", event.Body)

	str := event.Body
	b := Body{}
	err := json.Unmarshal([]byte(str), &b)
	if err != nil {
		fmt.Printf("Could not unmarshall: %v\n", event.Body)
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("Okay so your other function also executed successfully %v!", b.Name),
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
