package main

import (
	"context"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	operations "mealmates.com/lambda/DataBase/Operations"
)


type MyEvent struct {
	Recipe operations.RecipeJSON `json:"recipe"`
}

type MyResponse struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	operations.Put(request.Recipe)
	return MyResponse{}, nil
}

func main() {
	lambda.Start(HandleRequest)

	// HandleRequest(context.TODO(), MyEvent{})
}

// BUILD COMMAND:
// GOOS=linux go build main.go
