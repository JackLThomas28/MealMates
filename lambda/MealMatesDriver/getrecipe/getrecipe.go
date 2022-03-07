package getrecipe

import (
	"context"
	"encoding/json"
	"fmt"

	"mealmates.com/lambda/MealMatesDriver/mylambda"
	parseingredients "mealmates.com/lambda/MealMatesDriver/parseingredients"
)

type Response struct {
	Recipe Recipe `json:"recipe"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Rating struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type Nutrition struct {
	Calories string `json:"calories"`
}

type Recipe struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Image        Image     `json:"image"`
	Description  string    `json:"description"`
	PrepTime     string    `json:"prepTime"`
	CookTime     string    `json:"cookTime"`
	TotalTime    string    `json:"totalTime"`
	RecipeYield  string    `json:"recipeYield"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Categories   []string  `json:"categories"`
	Rating       Rating    `json:"rating"`
	Nutrition    Nutrition `json:"nutrition"`
	ParsedIngredients []parseingredients.Ingredient `json:"parsedIngredients"`
}

const (
	FILE_NAME = "getrecipe"
	LAMBDA_NAME = "getRecipe"
)

func GetRecipe(ctx context.Context, requestUrl string) (Response, error){
	const FUNC_NAME = " GetRecipe:"
	var resp Response

	inPayload := []byte(
		`{`                               +
			`"url": "` + requestUrl + `"` +
		`}`)
	
	outPayload, err := mylambda.InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
		return resp, err
	}
	// Convert the payload to a Response Object
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
	}
	return resp, err
}