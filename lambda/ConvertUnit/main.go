package main

import (
	"context"

	// Third Party
	lambda "github.com/aws/aws-lambda-go/lambda"
	convert "mealmates.com/lambda/ConvertUnit/Convert"
)

type MyEvent struct {
	Amount float64 `json:"amount"`
	To     string  `json:"to"`
	From   string  `json:"from"`
}

type MyResponse struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	result, err := convert.Convert(request.Amount).From(request.From).To(request.To)
	return MyResponse{Result: result}, err
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
