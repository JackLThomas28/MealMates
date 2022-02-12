package main

import (
	"context"
	"errors"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	ops "mealmates.com/lambda/DatabaseOps/ops"
	structs "mealmates.com/lambda/DatabaseOps/structs"
)


type MyEvent struct {
	Recipe structs.Recipe `json:"recipe"`
	Ingredients []structs.MyIngredient `json:"ingredients"`
	Operation string `json:"operation"`
	Table string `json:"table"`
}

type MyResponse struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	if request.Operation == "put" {
		ops.Put(ops.PutItem{Recipe: request.Recipe, Ingredients: request.Ingredients}, request.Table, request.BeginningIndex)
		return MyResponse{}, nil
	}
	return MyResponse{}, errors.New("Unrecognized operation: " + request.Operation)
}

func main() {
	lambda.Start(HandleRequest)

	// HandleRequest(context.TODO(), MyEvent{})
}

// BUILD COMMAND:
// GOOS=linux go build main.go
