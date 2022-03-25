package mylambda

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	// Third Party
	"github.com/JackLThomas28/MealMates/lambda/objects/ingredientdb"
	"github.com/JackLThomas28/MealMates/lambda/objects/recipe"
)

type Response struct {
	Success bool `json:"success"`
	Body []MyIngredient `json:"body"`
	RawPayload []byte `json:"rawpayload"`
}

type OperationError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorType string `json:"errorType"`
}

type ReqItem interface {
	IsRecipe() bool
	GetRecipe() MyRecipe
	IsIngredient() bool
	GetIngredient() MyIngredient
}

type MyIngredient ingredientdb.IngredientDB
func (i MyIngredient) IsIngredient() bool {
	return true
}
func (i MyIngredient) IsRecipe() bool {
	return false
}
func (i MyIngredient) GetIngredient() MyIngredient {
	return i
}
func (i MyIngredient) GetRecipe() MyRecipe {
	return MyRecipe{}
}

type MyRecipe recipe.Recipe
func (r MyRecipe) IsIngredient() bool {
	return false
}
func (r MyRecipe) IsRecipe() bool {
	return true
}
func (r MyRecipe) GetIngredient() MyIngredient {
	return MyIngredient{}
}
func (r MyRecipe) GetRecipe() MyRecipe {
	return r
}

const (
	file_name = "operations"
	lambda_name = "dataBaseOperations"
	// DB Operation
	GET = "get"
	UPDATE = "update"
	PUT = "put"
	// Table Names
	INGREDIENTS_TABLE = "Ingredients"
	RECIPES_TABLE = "AllRecipes"
)

func DatabaseOperation(ctx context.Context, reqObj ReqItem) (Response, error) {
	// Set the function name for error logging
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
		fmt.Println(file_name + FUNC_NAME, "invalid request object")
		// Error path
	}
	outPayload, err := InvokeLambda(ctx, lambda_name, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(file_name + FUNC_NAME, err)
		return resp, err
	}

	// Try converting the payload to a Response Object
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		// We didn't receive a normal Response...
		fmt.Println(file_name + FUNC_NAME, err)
	}
	resp.RawPayload = outPayload
	return resp, err
}

// Will add an ingredient or update an ingredient
func UpdateIngredientsTable(ctx context.Context, ingrs []MyIngredient, action string) error {
	var err error
	for _, ingr := range ingrs {
		inPayload := buildIngredientRequest(action, ingr)
		_, err = InvokeLambda(ctx, lambda_name, inPayload)
	}
	return err
}

// Builds an ingredient request item
func buildIngredientRequest(operation string, ingredient MyIngredient) []byte {
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
func buildRecipeRequest(operation string, recipe MyRecipe) []byte {
	// Request with ingredient
	body := []byte(
		`"recipe": {`                                 +
			`"id": "` + strconv.Itoa(recipe.ID) + `"` +
		`}`)

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
