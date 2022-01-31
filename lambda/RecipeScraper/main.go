package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	"mealmates.com/lambda/RecipeScraper/recipes"
	"mealmates.com/lambda/RecipeScraper/recipes/allrecipes"
	structs "mealmates.com/lambda/RecipeScraper/structs"
)

type MyEvent struct {
	URL string `json:"url"`
}

type MyResponse struct {
	Recipe structs.StandardRecipe `json:"recipe"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	// Currently only scraping "allrecipes.com" recipes
	allrecipesRecipe, err := allrecipes.GetRecipe(request.URL)
	if err != nil {
		return MyResponse{}, err
	}

	// Only save the information needed from the recipe
	standardRecipe, err := standard.StandardizeRecipe(allrecipesRecipe)
	if err != nil {
		return MyResponse{}, err
	}

	return MyResponse{Recipe: standardRecipe}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
