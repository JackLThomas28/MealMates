package main

import (
	"context"

	// Third Party
	"github.com/JackLThomas28/MealMates/lambda/objects/recipe"
	"github.com/aws/aws-lambda-go/lambda"

	// Local packages
	recipes "mealmates.com/lambda/RecipeScraper/recipes"
	allrecipes "mealmates.com/lambda/RecipeScraper/recipes/allrecipes"
)

type MyEvent struct {
	URL string `json:"url"`
}

type MyResponse struct {
	Recipe recipe.Recipe `json:"recipe"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	// Currently only scraping "allrecipes.com" recipes
	allrecipesRecipe, err := allrecipes.GetRecipe(request.URL)
	if err != nil {
		return MyResponse{}, err
	}

	// Only save the information needed from the recipe
	standardRecipe, err := recipes.StandardizeRecipe(allrecipesRecipe)
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
// zip main.zip main
// aws lambda update-function-code --function-name getRecipe --zip-file fileb://./main.zip