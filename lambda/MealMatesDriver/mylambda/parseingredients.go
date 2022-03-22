package mylambda

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JackLThomas28/MealMates/lambda/objects/ingredient"
)

func convertIngredientsToBytes(ingrs []string) ([]byte, error) {
	const (
		file_name = "parseingredients"
		lambda_name = "parseIngredients"
		func_name = " convertIngredientsToBytes:"
	)
	var err error

	var ingrInBytes []byte
	for i, ingr := range ingrs {
		tIngr, err := json.Marshal(ingr)
		if err != nil {
			fmt.Println(file_name + func_name, err)
		} else {
			ingrInBytes = append(ingrInBytes, tIngr...)
			// Don't add the last comma
			if i < len(ingrs) - 1 {
				ingrInBytes = append(ingrInBytes, byte(','))
			}
		}
	}

	return ingrInBytes, err
}

func ParseIngredients(ctx context.Context, ingrs []string) ([]ingredient.Ingredient, error) {
	const (
		file_name = "parseingredients"
		lambda_name = "parseIngredients"
		func_name = " ParseIngredients:"
	)

	// Convert the list of strings to bytes
	ingrsInbytes, err := convertIngredientsToBytes(ingrs)
	if err != nil {
		return []ingredient.Ingredient{}, err
	}

	inPayload := []byte(
		`{`                   +
			`"ingredients": [`)
	inPayload = append(inPayload, ingrsInbytes...)
	inPayload = append(inPayload, []byte(
			`]`               +
		`}`)...)

	outPayload, err := InvokeLambda(ctx, lambda_name, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(file_name + func_name, err)
		return []ingredient.Ingredient{}, err
	}

	// Convert the payload to a Response object
	type response struct {
		Ingredients []ingredient.Ingredient `json:"ingredients"`
	}
	var resp response
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		fmt.Println(file_name + func_name, err)
	}
	return resp.Ingredients, err
}