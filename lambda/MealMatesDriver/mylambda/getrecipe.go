package mylambda

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JackLThomas28/MealMates/lambda/objects/recipe"
)

const (
	FILE_NAME = "getrecipe"
	LAMBDA_NAME = "getRecipe"
)

func GetRecipe(ctx context.Context, requestUrl string) (recipe.Recipe, error){
	const (
		FILE_NAME = "getrecipe"
		LAMBDA_NAME = "getRecipe"
		FUNC_NAME = " GetRecipe:"
	)
	
	inPayload := []byte(
		`{`                           +
		`"url": "` + requestUrl + `"` +
		`}`)
		
	outPayload, err := InvokeLambda(ctx, LAMBDA_NAME, inPayload)
	// Only continue if no errors
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
		return recipe.Recipe{}, err
	}

	// Convert the payload to a Response Object
	type response struct {
		Recipe recipe.Recipe `json:"recipe"`
	}
	var resp response
	err = json.Unmarshal(outPayload, &resp)
	if err != nil {
		fmt.Println(FILE_NAME + FUNC_NAME, err)
	}

	return resp.Recipe, err
}