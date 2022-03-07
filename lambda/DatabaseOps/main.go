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
	Body []reqitem.RequestItem `json:"body"`
	// Body map[string]types.AttributeValue `json:"body"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {	
	myResponse := MyResponse{Success: true}
	var scanResults []reqitem.RequestItem
	var err error

	// Resolve the request item
	var reqItem reqitem.RequestItem
	if request.Table == reqitem.RECIPE_TABLE_NAME {
		reqItem = &request.Recipe
	} else if request.Table == reqitem.INGR_TABLE_NAME {
		reqItem = &request.Ingredient
	} else {
		myResponse.Success = false
		err = errors.New("main handlerequest: invalid table name: " + request.Table)
	}
	// Only continue if table exists
	if err != nil {
		return myResponse, err
	}

	// Determine which type of operation to perform and perform it
	if request.Operation == "put" {
		err = ops.Put(reqItem)
	} else if request.Operation == "update" {
		err = ops.Update(reqItem) 
	} else if request.Operation == "delete" {
		err = ops.Delete(reqItem)
	} else if request.Operation == "scan" {
		scanResults, err = ops.Scan(reqItem)
	} else if request.Operation == "get" {
		err = ops.Get(reqItem)
	} else {
		err = errors.New("main handlerequest: invalid db operation " + request.Operation)
	}
	// Only continue if valid db operation
	if err != nil {
		myResponse.Success = false
		return myResponse, err
	} else {
		// Append the results
		if len(scanResults) > 0 {
			myResponse.Body = scanResults
		} else {
			myResponse.Body = append(myResponse.Body, reqItem)
		}
	}
	return myResponse, err
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMANDS:
// GOOS=linux go build main.go
// zip main.zip main
// aws lambda update-function-code --function-name dataBaseOperations --zip-file fileb://./main.zip
