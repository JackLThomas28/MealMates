package allrecipes

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	// "log"

	// Local Packages
	myutils "mealmates.com/lambda/RecipeScraper/myutils"
	structs "mealmates.com/lambda/RecipeScraper/structs"
)

const URL = "https://www.allrecipes.com/recipe/"
const FILE_NAME = "allrecipes.json"

func GetRecipe(inputURL string) (structs.ARRecipe, error) {
	// Get the recipe ID from the URL
	// Ex: "https://www.allrecipes.com/recipe/279984/air-fryer-sweet-and-spicy-roasted-carrots/"
	re, _ := regexp.Compile(URL)
	indices := re.FindStringIndex(inputURL)
	// Only save the last part of the url
	// Ex: "279984/air-fryer-sweet-and-spicy-roasted-carrots/"
	recipeId := inputURL[indices[1]:]
	re, _ = regexp.Compile("/")
	indices = re.FindStringIndex(recipeId)
	// Only save the numbers
	// Ex: "279984"
	recipeId = recipeId[:indices[0]]
	// Convert from string to int
	recipeIdInt, err := strconv.Atoi(recipeId)
	if err != nil {
		return structs.ARRecipe{}, errors.New("Could not find/convert recipe ID")
	}

	node := myutils.GetHtmlNode(inputURL)

	const TYPE = "application/ld+json"
	n := myutils.GetElementByType(node, TYPE)
	if n == nil {
		return structs.ARRecipe{}, errors.New("Could not get element by type.")
	}

	// Store the data in Recipe struct
	var recipe []structs.ARRecipe
	json.Unmarshal([]byte(n.FirstChild.Data), &recipe)

	const RECIPEINDEX = 1
	if recipe[RECIPEINDEX].MainEntityOfPage == "" {
		return structs.ARRecipe{}, errors.New("Could not get MainEntityOfPage property.")
	}

	// log.Printf("AllRecipes")
	// for i,_ := range recipe[RECIPEINDEX].RecipeIngredient {
	// 	log.Printf(recipe[RECIPEINDEX].RecipeIngredient[i])
	// }
	recipe[RECIPEINDEX].ID = recipeIdInt
	return recipe[RECIPEINDEX], nil
}
