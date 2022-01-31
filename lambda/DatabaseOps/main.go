package main

import (
	"context"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	structs "mealmates.com/lambda/DatabaseOps/structs"
	ops "mealmates.com/lambda/DatabaseOps/ops"
)


type MyEvent struct {
	Recipe structs.Recipe `json:"recipe"`
}

type MyResponse struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	ops.Put(request.Recipe)
	return MyResponse{}, nil
}

func main() {
	lambda.Start(HandleRequest)

	// HandleRequest(context.TODO(), MyEvent{})
}

// BUILD COMMAND:
// GOOS=linux go build main.go
