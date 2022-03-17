package databaseops

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	// Local
	"mealmates.com/lambda/MealMatesDriver/mylambda"
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

type ReqItem interface {
	GetRecipeIds() []int
}

type Ingredient struct {
	Name string `json:"name"`
	RecipeIds []int `json:"recipeIds"`
	// TODO: Remove RecipeId
	RecipeId int `json:"recipeId"`
}

func (i Ingredient)GetRecipeIds() []int {
	return i.RecipeIds
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
)

func DatabaseOperation(ctx context.Context, ingr Ingredient) (Response, error) {
	const FUNC_NAME = " DatabaseOperation:"
	var resp Response

	inPayload := buildIngredientRequest(GET, INGREDIENTS_TABLE, ingr)
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
		inPayload := buildIngredientRequest(action, INGREDIENTS_TABLE, ingr)
		_, err = mylambda.InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	}
	return err
}

// Builds an ingredient request item
func buildIngredientRequest(operation string, table string, ingredient Ingredient) []byte {
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
