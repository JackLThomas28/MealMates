package standard

import (
	"mealmates.com/lambda/RecipeScraper/recipes/allrecipes"
	// "encoding/json"
	// "errors"
)

type ImageJSON struct {
	URL    string `json:"URL"`
	Width  int    `json:"Width"`
	Height int    `json:"Height"`
}

type RatingJSON struct {
	RatingValue float32 `json:"RatingValue"`
	RatingCount int     `json:"RatingCount"`
}

type NutritionJSON struct {
	Calories string `json:"Calories"`
}

type RecipeJSON struct {
	ID           int           `json:"ID"`
	Name         string        `json:"Name"`
	Image        ImageJSON     `json:"Image"`
	Description  string        `json:"Description"`
	PrepTime     string        `json:"PrepTime"`
	CookTime     string        `json:"CookTime"`
	TotalTime    string        `json:"TotalTime"`
	RecipeYield  string        `json:"RecipeYield"`
	Ingredients  []string      `json:"Ingredients"`
	Instructions []string      `json:"Instructions"`
	Categories   []string      `json:"Categories"`
	Rating       RatingJSON    `json:"Rating"`
	Nutrition    NutritionJSON `json:"Nutrition"`
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
