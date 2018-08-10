package main

import (
	"context"

	_ "github.com/hexoul/ether-stealer/crypto"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := "WRONG"
	statusCode := 200

	// Validate privkey
	if method := request.QueryStringParameters["key"]; method != "" {
		respBody = "GOOD"
	}

	return events.APIGatewayProxyResponse{Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
