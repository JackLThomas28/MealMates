package main

import (
	"context"

	// Local Packages
	"github.com/aws/aws-lambda-go/lambda"
	parser "mealmates.com/lambda/IngredientParser/Parser"
)

type MyEvent struct {
	Ingredients []string `json:"ingredients"`
}

type MyResponse struct {
	Ingredients []parser.Ingredient `json:"ingredients"`
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