package mylambda

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/JackLThomas28/MealMates/lambda/objects/recipeandfrequency"
)

type RF recipeandfrequency.RecipeAndFrequency

func RunAlgorithm(ctx context.Context, ingredientDBs []MyIngredient, cntItemsToReturn int) ([]RF, error){
	const (
		file_name = "runalgorithm"
		lambda_name = "runAlgorithm"
		func_name = " RunAlgorithm:"
	)

	ingrsInBytes, err := convertIngredientDBsToBytes(ingredientDBs)

	inPayload := []byte(
		`{` +
			`"cntItemsToReturn": ` + strconv.Itoa(cntItemsToReturn) + `,` +
			`"ingredients": [`)
	inPayload = append(inPayload, ingrsInBytes...)
	inPayload = append(inPayload, []byte(
			`]` +
		`}`)...)

	outPayload, err := InvokeLambda(ctx, lambda_name, inPayload)
	if err != nil {
		fmt.Println(file_name + func_name, err)
		// return
	}

	type response struct {
		RF []RF `json:"recipes"`
	}
	var resp response
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		fmt.Println(file_name + func_name, err)
	}
	return resp.RF, err
}

func convertIngredientDBsToBytes(ingrs []MyIngredient) ([]byte, error) {
	const (
		file_name = "runalgorithm"
		lambda_name = "runAlgorithm"
		func_name = " convertIngredientDBsToBytes:"
	)
	var err error
	var ingrsInBytes []byte
	
	for i, ingr := range ingrs {
		// Get the ingredient recipe ids in bytes
		var idsInBytes []byte
		for j, id := range ingr.RecipeIds {
			idsInBytes = append(idsInBytes, []byte(strconv.Itoa(id))...)
			
			// Don't append ',' for last id
			if j < len(ingr.RecipeIds) - 1 {
				idsInBytes = append(idsInBytes, byte(','))
			}
		}

		bIngr := []byte(
			`{` +
				`"name": "` + ingr.Name + `",` +
				`"recipeIds": [`)
		bIngr = append(bIngr, idsInBytes...)
		bIngr = append(bIngr, []byte(
				`]` + 
			`}`)...)
		
		// Don't append ',' for last ingr
		if i < len(ingrs) - 1 {
			bIngr = append(bIngr, byte(','))
		}
		ingrsInBytes = append(ingrsInBytes, bIngr...)
	}
	return ingrsInBytes, err
}