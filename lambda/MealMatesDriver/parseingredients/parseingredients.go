package parseingredients

import (
	"context"
	"encoding/json"
	"fmt"

	// Local
	"mealmates.com/lambda/MealMatesDriver/mylambda"
)

const (
	FILE_NAME = "parseingredients"
	LAMBDA_NAME = "parseIngredients"
)

type Ingredient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
	Raw    string  `json:"raw"`
}

type Response struct {
	Ingredients []Ingredient `json:"ingredients"`
}

func convertIngredientsToBytes(ingrs []string) ([]byte, error) {
	const FUNC_NAME = " convertIngredientsToBytes:"
	var err error

	var ingrInBytes []byte
	for i, ingr := range ingrs {
		tIngr, err := json.Marshal(ingr)
		if err != nil {
			fmt.Println(FILE_NAME + FUNC_NAME, err)
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

func ParseIngredients(ctx context.Context, ingrs []string) (Response, error) {
	const FUNC_NAME = " ParseIngredients:"
	var resp Response

	// Convert the list of strings to bytes
	ingrsInbytes, err := convertIngredientsToBytes(ingrs)
	if err != nil {
		return resp, err
	}

	inPayload := []byte(
		`{`                   +
			`"ingredients": [`)
	inPayload = append(inPayload, ingrsInbytes...)
	inPayload = append(inPayload, []byte(
			`]`               +
		`}`)...)

	outPayload, err := mylambda.InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
		return resp, err
	}
	// Convert the payload to a Response object
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
	}
	return resp, err
}