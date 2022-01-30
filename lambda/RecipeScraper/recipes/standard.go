package standard

import (
	"mealmates.com/lambda/RecipeScraper/recipes/allrecipes"
	// "encoding/json"
	// "errors"
)

type ImageJSON struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type RatingJSON struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type NutritionJSON struct {
	Calories string `json:"calories"`
}

type RecipeJSON struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Image        ImageJSON     `json:"image"`
	Description  string        `json:"description"`
	PrepTime     string        `json:"prepTime"`
	CookTime     string        `json:"cookTime"`
	TotalTime    string        `json:"totalTime"`
	RecipeYield  string        `json:"recipeYield"`
	Ingredients  []string      `json:"ingredients"`
	Instructions []string      `json:"instructions"`
	Categories   []string      `json:"categories"`
	Rating       RatingJSON    `json:"rating"`
	Nutrition    NutritionJSON `json:"nutrition"`
}

func StandardizeRecipe(recipe allrecipes.Recipe) (RecipeJSON, error) {
	// Build the recipe
	var newRecipe RecipeJSON
	newRecipe.ID = recipe.ID
	newRecipe.Name = recipe.Name

	var newImage ImageJSON
	newImage.URL = recipe.Image.URL
	newImage.Width = recipe.Image.Width
	newImage.Height = recipe.Image.Height

	newRecipe.Image = newImage
	newRecipe.Description = recipe.Description
	newRecipe.PrepTime = recipe.PrepTime
	newRecipe.CookTime = recipe.CookTime
	newRecipe.TotalTime = recipe.TotalTime
	newRecipe.RecipeYield = recipe.RecipeYield
	for i := 0; i < len(recipe.RecipeIngredient); i++ {
		newRecipe.Ingredients = append(newRecipe.Ingredients, recipe.RecipeIngredient[i])
	}
	for i := 0; i < len(recipe.RecipeInstructions); i++ {
		newRecipe.Instructions = append(newRecipe.Instructions, recipe.RecipeInstructions[i].Text)
	}
	for i := 0; i < len(recipe.RecipeCategory); i++ {
		newRecipe.Categories = append(newRecipe.Categories, recipe.RecipeCategory[i])
	}

	var newRating RatingJSON
	newRating.RatingValue = recipe.AggregateRating.RatingValue
	newRating.RatingCount = recipe.AggregateRating.RatingCount

	newRecipe.Rating = newRating

	var newNutrition NutritionJSON
	newNutrition.Calories = recipe.Nutrition.Calories
	newRecipe.Nutrition = newNutrition

	return newRecipe, nil
}
