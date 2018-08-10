package main

import (
	"context"

	"github.com/hexoul/ether-stealer/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := "WRONG"
	statusCode := 200

	// Validate privkey
	if key := request.QueryStringParameters["key"]; key != "" {
		privkey := hexutil.MustDecode(key)
		if addr, err := crypto.ToAddressFromPrivkey(privkey); err == nil {
			respBody = hexutil.Encode(addr.Bytes())
		}
	}

	return events.APIGatewayProxyResponse{Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
