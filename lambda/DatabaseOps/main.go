package main

import (
	"context"
	"errors"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local
	ingredient "mealmates.com/lambda/DatabaseOps/ingredient"
	ops "mealmates.com/lambda/DatabaseOps/ops"
	recipe "mealmates.com/lambda/DatabaseOps/recipe"
)

type MyEvent struct {
	Recipe     recipe.Recipe         `json:"recipe"`
	Ingredient ingredient.Ingredient `json:"ingredient"`
	Operation  string                `json:"operation"`
	Table      string                `json:"table"`
}

type MyResponse struct {
	Success bool `json:"success"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	var err error
	// Determine which type of operation to perform and perform it
	if request.Operation == "put" {
		err = ops.Put(ops.PutItem{Recipe: request.Recipe, Ingredient: request.Ingredient}, request.Table)
	} else if request.Operation == "update" {
		err = ops.Update(ops.UpdateItem{Recipe: request.Recipe, Ingredient: request.Ingredient}, request.Table) 
	} else if request.Operation == "delete" {
		err = ops.Delete(ops.DeleteItem{Recipe: request.Recipe, Ingredient: request.Ingredient}, request.Table)
	} else if request.Operation == "scan" {
		_, err = ops.Scan(ops.ScanItem{Recipe: request.Recipe, Ingredient: request.Ingredient}, request.Table)
	} else if request.Operation == "get" {
		_, err = ops.Get(ops.GetItem{Recipe: request.Recipe, Ingredient: request.Ingredient}, request.Table)
	} else {
		err = errors.New("Error: invalid db operation " + request.Operation)
	}

	if err != nil {
		return MyResponse{Success: false}, err
	}
	return MyResponse{Success: true}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
