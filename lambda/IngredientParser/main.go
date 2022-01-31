package main

import (
	"context"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local Packages
	parser "mealmates.com/lambda/IngredientParser/parser"
	structs "mealmates.com/lambda/IngredientParser/structs"
)

type MyEvent struct {
	Ingredients []string `json:"ingredients"`
}

type MyResponse struct {
	Ingredients []structs.Ingredient `json:"ingredients"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	ingredients, err := parser.ParseIngredients(request.Ingredients)
	return MyResponse{Ingredients: ingredients}, err
}

func main() {
	lambda.Start(HandleRequest)

	// var testInput []string
	// testInput = append(testInput, "¾ cup chopped green bell pepper")
	// ingredients, err := parser.ParseIngredients(testInput)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(ingredients)
	// }
}

// BUILD COMMAND:
// GOOS=linux go build main.go