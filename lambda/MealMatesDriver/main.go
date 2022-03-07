package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local
	"mealmates.com/lambda/MealMatesDriver/databaseoperations"
	"mealmates.com/lambda/MealMatesDriver/getrecipe"
	"mealmates.com/lambda/MealMatesDriver/mylambda"
	"mealmates.com/lambda/MealMatesDriver/parseingredients"
)

type MyRequest struct {
	URL string `json:"url"`
}

type MyResponse struct {
	// Recipe structs.StandardRecipe `json:"recipe"`
}


func GetDatabaseRequestInBytes(operation string, table string, ingredient parseingredients.DBIngredient, recipe getrecipe.Recipe) []byte {
	var body []byte
	// Request with ingredient
	if ingredient.Name != "" && recipe.Name == "" {
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
	} else if ingredient.Name == "" && recipe.Name != "" {
		// TODO: Implement when we recieve a recipe
	} else {
		// TODO: Implement when we don't recieve ingredient or recipe
	}

	// Build the request
	request := []byte(
		`{`                                     +
			`"operation": "` + operation + `",` +
			`"table": "Ingredients",`)
	request = append(request, body...)
	request = append(request, []byte(
		`}`)...)
	
	return request
}

func HandleRequest(ctx context.Context, request MyRequest) (MyResponse, error) {
	myResponse := MyResponse{}

	// ---------------------------------------------------------------------- //
	// 1. Get the Recipe
	// ---------------------------------------------------------------------- //

	grResp, err := getrecipe.GetRecipe(ctx, request.URL)
	// payload, err := mylambda.InvokeLambda(ctx, "getRecipe", []byte(`{"url":"` + request.URL + `"}`))
	// // Only continue if no errors
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return myResponse, err
	// }
	// // Convert the payload to a Recipe object
	// var gr getrecipe.GetRecipeResponse
	// err = json.Unmarshal(payload, &gr)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return myResponse, err
	// }

	// ---------------------------------------------------------------------- //
	// 2. Parse the Ingredients
	// ---------------------------------------------------------------------- //

	piResp, err := parseingredients.ParseIngredients(ctx, grResp.Recipe.Ingredients)
	// Convert the ingredients to []byte
	// ingrListBytes := []byte(`{"ingredients": [`)
	// for i, ingr := range gr.Recipe.Ingredients {
	// 	tIngr, err := json.Marshal(ingr)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		ingrListBytes = append(ingrListBytes, tIngr...)
	// 		// Don't add the last comma
	// 		if i < len(gr.Recipe.Ingredients) - 1 {
	// 			ingrListBytes = append(ingrListBytes, byte(','))
	// 		}
	// 	}
	// }
	// ingrListBytes = append(ingrListBytes, []byte(`]}`)...)

	// payload, err = mylambda.InvokeLambda(ctx, "parseIngredients", ingrListBytes)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return myResponse, err
	// }

	// ---------------------------------------------------------------------- //
	// 3. Scan the DB for the ingredients
	// ---------------------------------------------------------------------- //

	// Convert the payload to Ingredient objects
	var parseIngredients parseingredients.ParseIngredientsResponse
	err = json.Unmarshal(payload, &parseIngredients)
	if err != nil {
		fmt.Println("Error:", err)
		return myResponse, err
	}

	// Loop through each ingredient and query ingredient db.

	// List of ingredients that should be added to the db.
	var ingredientsToAdd []parseingredients.DBIngredient
	// List of ingredients to update
	var ingredientsToUpdate []parseingredients.DBIngredient
	// List of all ingredients in DB form
	var ingredients []parseingredients.DBIngredient

	// To Do: Rework the logic here. Results are dependent on Ingredient Parsing
	for _, pIngr := range parseIngredients.Ingredients {
		dbIngr := parseingredients.DBIngredient{Name: pIngr.Name}

		// Get the ingredient
		pIngrListBytes := GetDatabaseRequestInBytes("get", "Ingredients", dbIngr, getrecipe.Recipe{})
		payload, err = mylambda.InvokeLambda(ctx, "dataBaseOperations", pIngrListBytes)
		if err != nil {
			fmt.Println("Error:", err)
			return myResponse, err
		}
		fmt.Println("DB Payload:", string(payload))
		
		// ------------------------------------------------------------------ //
		// 3a. Add the ingredient to DB if it does not exist
		// ------------------------------------------------------------------ //

		// Check if we received an error
		var errResp databaseoperations.OperationError
		err = json.Unmarshal(payload, &errResp)
		if err != nil {
			fmt.Println("main handlerequest:", err)
		} else {
			// If we were able to unmarshal response into error -> 
			// then we received error response
			if errResp.ErrorMessage == "ingredient parseresult: could not locate ingredient" {
				fmt.Println("Could not locate ingredient " + pIngr.Name)
				dbIngredient := parseingredients.DBIngredient{Name: pIngr.Name, RecipeIds: []int{gr.Recipe.ID}}
				ingredientsToAdd = append(ingredientsToAdd, dbIngredient)

				// Append the ingredient to the master list
				ingredients = append(ingredients, dbIngredient)
			}
		}

		// Convert DB response to DB object
		var opResp databaseoperations.OperationResponse
		err = json.Unmarshal(payload, &opResp)
		if err != nil {
			fmt.Println("main handlerequest:", err)
		} else {
			// Check the ingredient's recipeIds
			found := false
			for _, item := range opResp.Body {
				if found {
					break
				}
				for _, id := range item.GetRecipeIds() {
					// Exit if we've found a match
					if id == gr.Recipe.ID {
						found = true
						break
					}
				}
			}

			// If we didn't find it, add it to the db
			if !found {
				dbIngredient := parseingredients.DBIngredient{Name: pIngr.Name, RecipeIds: []int{gr.Recipe.ID}}
				ingredientsToUpdate = append(ingredientsToUpdate, dbIngredient)

				// Append the ingredient to the master list
				// dbIngredient.RecipeIds = append(dbIngredient.RecipeIds, opResp.Body)
				// ingredients = append(ingredients, )
			}
		}

		// ------------------------------------------------------------------ //
		// 4. Determine which recipes share the most ingredients with original recipe
		// ------------------------------------------------------------------ //


	}

	// Add the missing ingredients to the db
	for _, dbIngr := range ingredientsToAdd {
		payloadIngrBytes := GetDatabaseRequestInBytes("put", "Ingredients", dbIngr, getrecipe.Recipe{})
		payload, err = mylambda.InvokeLambda(ctx, "dataBaseOperations", payloadIngrBytes)
	}

	// Update the ingredient's recipeIds list
	for _, dbIngr := range ingredientsToUpdate {
		payloadIngrBytes := GetDatabaseRequestInBytes("update", "Ingredients", dbIngr, getrecipe.Recipe{})
		payload, err = mylambda.InvokeLambda(ctx, "dataBaseOperations", payloadIngrBytes)
	}

	return myResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
// zip main.zip main
// aws lambda update-function-code --function-name mealmatesDriver --zip-file fileb://./main.zip
