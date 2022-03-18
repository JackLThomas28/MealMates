package databaseops

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	// Local
	"mealmates.com/lambda/MealMatesDriver/getrecipe"
	"mealmates.com/lambda/MealMatesDriver/mylambda"
	"mealmates.com/lambda/MealMatesDriver/parseingredients"
)

type Response struct {
	Success bool `json:"success"`
	Body []Ingredient `json:"body"`
	RawPayload []byte `json:"rawpayload"`
}

type OperationError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorType string `json:"errorType"`
}

type Ingredient struct {
	Name string `json:"name"`
	RecipeIds []int `json:"recipeIds"`
	// TODO: Remove RecipeId
	RecipeId int `json:"recipeId"`
}

type Recipe struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Image        getrecipe.Image     `json:"image"`
	Description  string    `json:"description"`
	PrepTime     string    `json:"prepTime"`
	CookTime     string    `json:"cookTime"`
	TotalTime    string    `json:"totalTime"`
	RecipeYield  string    `json:"recipeYield"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Categories   []string  `json:"categories"`
	Rating       getrecipe.Rating    `json:"rating"`
	Nutrition    getrecipe.Nutrition `json:"nutrition"`
	ParsedIngredients []parseingredients.Ingredient `json:"parsedIngredients"`
}

type ReqItem interface {
	IsRecipe() bool
	GetRecipe() Recipe
	IsIngredient() bool
	GetIngredient() Ingredient
}
func (i Ingredient) IsIngredient() bool {
	return true
}
func (i Ingredient) IsRecipe() bool {
	return false
}
func (i Ingredient) GetIngredient() Ingredient {
	return i
}
func (i Ingredient) GetRecipe() Recipe {
	return Recipe{}
}

func (r Recipe) IsIngredient() bool {
	return false
}
func (r Recipe) IsRecipe() bool {
	return true
}
func (r Recipe) GetIngredient() Ingredient {
	return Ingredient{}
}
func (r Recipe) GetRecipe() Recipe {
	return r
}

const (
	FILE_NAME = "operations"
	LAMBDA_NAME = "dataBaseOperations"
	// DB Operation
	GET = "get"
	UPDATE = "update"
	PUT = "put"
	// Table Names
	INGREDIENTS_TABLE = "Ingredients"
	RECIPES_TABLE = "AllRecipes"
)

func DatabaseOperation(ctx context.Context, reqObj ReqItem) (Response, error) {
	const FUNC_NAME = " DatabaseOperation:"
	var resp Response
	var inPayload []byte

	if reqObj.IsRecipe() {
		recipe := reqObj.GetRecipe()
		inPayload = buildRecipeRequest(GET, recipe)

	} else if reqObj.IsIngredient() {
		ingr := reqObj.GetIngredient()
		inPayload = buildIngredientRequest(GET, ingr)
	} else {
		// Error path
	}
	outPayload, err := mylambda.InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
		return resp, err
	}

	// Try converting the payload to a Response Object
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		// We didn't receive a normal Response...
		fmt.Println(FILE_NAME + FUNC_NAME, err)
	}
	resp.RawPayload = outPayload
	return resp, err
}

// Will add an ingredient or update an ingredient
func UpdateIngredientsTable(ctx context.Context, ingrs []Ingredient, action string) error {
	var err error
	for _, ingr := range ingrs {
		inPayload := buildIngredientRequest(action, ingr)
		_, err = mylambda.InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	}
	return err
}

// Builds an ingredient request item
func buildIngredientRequest(operation string, ingredient Ingredient) []byte {
	var body []byte
	// Request with ingredient
	if ingredient.Name != "" {
		// Convert the list of recipeIds to bytes
		var idsBytes []byte
		for i, id := range ingredient.RecipeIds {
			idStr := strconv.Itoa(id)
			idsBytes = append(idsBytes, []byte(idStr)...)
			// Don't append ',' to last recipe Id in list
			if i < len(ingredient.RecipeIds) - 1 {
				idsBytes = append(idsBytes, byte(','))
			}
		}

		body = []byte(
			`"ingredient": {`                        +
				`"name": "` + ingredient.Name + `",` +
				`"recipeIds": [`)
		body = append(body, idsBytes...)
		body = append(body, []byte(
				`]`                                  +
			`}`)...)
	} else {
		// TODO: Implement when we don't recieve ingredient
	}
	return getRequestObject(operation, body, INGREDIENTS_TABLE)
}

// Builds a recipe request item
func buildRecipeRequest(operation string, recipe Recipe) []byte {
	var body []byte
	// Request with ingredient
	if recipe.Name != "" {
		body = []byte(
			`"recipe": {`                                 +
				`"id": "` + strconv.Itoa(recipe.ID) + `"` +
			`}`)
	} else {
		// TODO: Implement when we don't recieve ingredient
	}

	return getRequestObject(operation, body, RECIPES_TABLE)
}

func getRequestObject(operation string, body []byte, tableName string) []byte {
	request := []byte(
		`{`                                     +
			`"operation": "` + operation + `",` +
			`"table": "` + tableName + `",`)
	request = append(request, body...)
	request = append(request, []byte(
		`}`)...)

	return request
}
