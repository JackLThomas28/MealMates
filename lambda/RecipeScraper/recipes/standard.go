package standard

import (
	// "encoding/json"
	// "errors"

	// Local Packages
	structs "mealmates.com/lambda/RecipeScraper/structs"
)

func StandardizeRecipe(recipe structs.ARRecipe) (structs.StandardRecipe, error) {
	// Build the recipe
	var newRecipe structs.StandardRecipe
	newRecipe.ID = recipe.ID
	newRecipe.Name = recipe.Name

	var newImage structs.StandardImage
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

	var newRating structs.StandardRating
	newRating.RatingValue = recipe.AggregateRating.RatingValue
	newRating.RatingCount = recipe.AggregateRating.RatingCount

	newRecipe.Rating = newRating

	var newNutrition structs.StandardNutrition
	newNutrition.Calories = recipe.Nutrition.Calories
	newRecipe.Nutrition = newNutrition

	return newRecipe, nil
}
