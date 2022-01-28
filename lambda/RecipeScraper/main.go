package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	"mealmates.com/lambda/RecipeScraper/recipes"
	"mealmates.com/lambda/RecipeScraper/recipes/allrecipes"
)

type MyEvent struct {
	URL string `json:"url"`
}

func HandleRequest(ctx context.Context, request MyEvent) (standard.RecipeJSON, error) {
	// Currently only scraping "allrecipes.com" recipes
	allrecipesRecipe, err := allrecipes.GetRecipe(request.URL)
	if err != nil {
		return standard.RecipeJSON{}, err
	}

	// Only save the information needed by standardizing the recipe
	standardRecipe, err := standard.StandardizeRecipe(allrecipesRecipe)
	if err != nil {
		return standard.RecipeJSON{}, err
	}

	return standardRecipe, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
