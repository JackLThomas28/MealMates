package main

import (
	"context"

	// Third Party
	"github.com/JackLThomas28/MealMates/lambda/objects/ingredient"
	"github.com/aws/aws-lambda-go/lambda"

	// Local Packages
	parser "mealmates.com/lambda/IngredientParser/parser"
)

type MyEvent struct {
	Ingredients []string `json:"ingredients"`
}

type MyResponse struct {
	Ingredients []ingredient.Ingredient `json:"ingredients"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	ingredients, err := parser.ParseIngredients(request.Ingredients)
	return MyResponse{Ingredients: ingredients}, err
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
// zip main.zip main
// aws lambda update-function-code --function-name parseIngredients --zip-file fileb://./main.zip