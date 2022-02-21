package main

import (
	"context"
	"errors"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local

	ops "mealmates.com/lambda/DatabaseOps/ops"
	"mealmates.com/lambda/DatabaseOps/reqitem"
)

type MyEvent struct {
	Recipe     reqitem.Recipe     `json:"recipe"`
	Ingredient reqitem.Ingredient `json:"ingredient"`
	Operation  string             `json:"operation"`
	Table      string             `json:"table"`
}

type MyResponse struct {
	Success bool `json:"success"`
	Body reqitem.RequestItem `json:"body"`
	// Body map[string]types.AttributeValue `json:"body"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {	
	var err error

	// Resolve the request item
	var reqItem reqitem.RequestItem
	if request.Table == reqitem.RECIPE_TABLE_NAME {
		reqItem = &request.Recipe
	} else if request.Table == reqitem.INGR_TABLE_NAME {
		reqItem = &request.Ingredient
	} else {
		err = errors.New("Error: invalid table name: " + request.Table)
	}
	// Only continue if table exists
	if err != nil {
		return MyResponse{Success: false}, err
	}

	// Determine which type of operation to perform and perform it
	// var output reqitem.RequestItem
	// var output map[string]types.AttributeValue
	// var scanOutput *dynamodb.ScanOutput
	if request.Operation == "put" {
		err = ops.Put(reqItem)
	} else if request.Operation == "update" {
		err = ops.Update(reqItem) 
	} else if request.Operation == "delete" {
		err = ops.Delete(reqItem)
	} else if request.Operation == "scan" {
		_, err = ops.Scan(reqItem)
	} else if request.Operation == "get" {
		_, err = ops.Get(reqItem)
	} else {
		err = errors.New("Error: invalid db operation " + request.Operation)
	}
	// Only continue if valid db operation
	if err != nil {
		return MyResponse{Success: false}, err
	}
	return MyResponse{Success: true, Body: reqItem}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
