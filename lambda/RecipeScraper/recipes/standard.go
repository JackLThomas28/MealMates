package standard

import (
	// "encoding/json"
	// "errors"

	// Third Party
	"github.com/JackLThomas28/MealMates/lambda/objects/recipe"

	// Local Packages
	structs "mealmates.com/lambda/RecipeScraper/structs"
)

func StandardizeRecipe(rec structs.ARRecipe) (recipe.Recipe, error) {
	// Build the recipe
	var newRecipe recipe.Recipe
	newRecipe.ID = rec.ID
	newRecipe.Name = rec.Name

	var newImage recipe.Image
	newImage.URL = rec.Image.URL
	newImage.Width = rec.Image.Width
	newImage.Height = rec.Image.Height

	newRecipe.Image = newImage
	newRecipe.Description = rec.Description
	newRecipe.PrepTime = rec.PrepTime
	newRecipe.CookTime = rec.CookTime
	newRecipe.TotalTime = rec.TotalTime
	newRecipe.RecipeYield = rec.RecipeYield
	for i := 0; i < len(rec.RecipeIngredient); i++ {
		newRecipe.Ingredients = append(newRecipe.Ingredients, rec.RecipeIngredient[i])
	}
	for i := 0; i < len(rec.RecipeInstructions); i++ {
		newRecipe.Instructions = append(newRecipe.Instructions, rec.RecipeInstructions[i].Text)
	}
	for i := 0; i < len(rec.RecipeCategory); i++ {
		newRecipe.Categories = append(newRecipe.Categories, rec.RecipeCategory[i])
	}

	var newRating recipe.Rating
	newRating.RatingValue = rec.AggregateRating.RatingValue
	newRating.RatingCount = rec.AggregateRating.RatingCount

	newRecipe.Rating = newRating

	var newNutrition recipe.Nutrition
	newNutrition.Calories = rec.Nutrition.Calories
	newRecipe.Nutrition = newNutrition

	return newRecipe, nil
}
