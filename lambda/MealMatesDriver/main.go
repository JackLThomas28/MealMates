package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"

	// Third Party
	"github.com/aws/aws-lambda-go/lambda"

	// Local
	databaseops "mealmates.com/lambda/MealMatesDriver/databaseoperations"
	"mealmates.com/lambda/MealMatesDriver/getrecipe"
	"mealmates.com/lambda/MealMatesDriver/mylambda"
	"mealmates.com/lambda/MealMatesDriver/parseingredients"
)

type MyRequest struct {
	URL string `json:"url"`
}

type MyResponse struct {
	Recipes []getrecipe.Recipe `json:"recipes"`
}

const (
	FILE_NAME = "main"
)

func HandleRequest(ctx context.Context, request MyRequest) (MyResponse, error) {
	const FUNC_NAME = " HandleRequest: "
	myResponse := MyResponse{}

	// ---------------------------------------------------------------------- //
	// 1. Get the Recipe
	// ---------------------------------------------------------------------- //

	grResp, err := getrecipe.GetRecipe(ctx, request.URL)
	if err != nil {
		return myResponse, err
	}
	fmt.Println("GetRecipe Response", grResp)
	dbRecipe := databaseops.Recipe{ID: grResp.Recipe.ID}

	// ---------------------------------------------------------------------- //
	// 2. Parse the Ingredients
	// ---------------------------------------------------------------------- //

	piResp, err := parseingredients.ParseIngredients(ctx, grResp.Recipe.Ingredients)
	if err != nil {
		return myResponse, err
	}
	fmt.Println("ParseIngredient Response", piResp)

	// ---------------------------------------------------------------------- //
	// 3. Add the Recipe to the DB if it isn't already there
	// ---------------------------------------------------------------------- //
	doResp, err := databaseops.DatabaseOperation(ctx, dbRecipe)
	if err != nil {
		return myResponse, err
	}
	fmt.Println("DatabaseOperation Response", doResp)
	
	// ---------------------------------------------------------------------- //
	// 4. Scan the DB for the ingredients
	// ---------------------------------------------------------------------- //
	// Loop through each ingredient and query ingredient db.

	// List of ingredients that should be added to the db.
	var ingredientsToAdd []databaseops.Ingredient
	// List of ingredients to update
	var ingredientsToUpdate []databaseops.Ingredient
	// List of all ingredients in DB form
	var allIngredients []databaseops.Ingredient

	// To Do: Rework the logic here. Results are dependent on Ingredient Parsing
	for _, pIngr := range piResp.Ingredients {
		dIngr := databaseops.Ingredient{Name: pIngr.Name, RecipeIds: []int{grResp.Recipe.ID}}
		doResp, err = databaseops.DatabaseOperation(ctx, dIngr)
		fmt.Println("DatabaseOperation Response", doResp)
		// Move this logic to DB Driver
		if err != nil || !doResp.Success {
			// Try converting the payload to an Error Response Object
			var errRsp mylambda.ErrorResponse
			err = json.Unmarshal(doResp.RawPayload, &errRsp)
			if err != nil {
				// We didn't receive an error response either...exit
				fmt.Println(FILE_NAME + FUNC_NAME, err)
				return myResponse, err
			} else {
				// We received an error - likely ingredient is missing from db
				if errRsp.ErrorMessage != "ingredient parseresult: could not locate ingredient" {
					return myResponse, errors.New(FILE_NAME + FUNC_NAME + "unknown error")
				}

				fmt.Println("Could not located ingredient " + dIngr.Name)
				fmt.Println("Adding ingredient", dIngr)
				ingredientsToAdd = append(ingredientsToAdd, dIngr)
			}
		}
		// Check the ingredient to see if it contains this recipe's ID
		found := false
		for _, item := range doResp.Body {
			dIngr = item
			if found {
				break
			}
			for _, id := range item.RecipeIds {
				// Exit if we've found a match
				if id == grResp.Recipe.ID {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Ingredient " + dIngr.Name + " does not contain recipe ID " + strconv.Itoa(dIngr.RecipeIds[0]))
				fmt.Println("Updating ingredient", dIngr)
				ingredientsToUpdate = append(ingredientsToUpdate, dIngr)
			}
		}
		// Add the ingredient to the master list
		allIngredients = append(allIngredients, dIngr)
	}

	databaseops.UpdateIngredientsTable(ctx, ingredientsToAdd, databaseops.PUT)
	databaseops.UpdateIngredientsTable(ctx, ingredientsToUpdate, databaseops.UPDATE)

	// ---------------------------------------------------------------------- //
	// 5. Determine which recipes share the most ingredients with original recipe
	// ---------------------------------------------------------------------- //

	fmt.Println(allIngredients)
	recipeIdAmts := make(map[int]int)
	for _, ingr := range allIngredients {
		for _, id := range ingr.RecipeIds {
			if recipeIdAmts[id] == 0 {
				recipeIdAmts[id] = 1
			} else {
				recipeIdAmts[id] += 1
			}
		}
	}

	var recipesAndOccurrences []RecipeAndOccurrence
	for key, val := range recipeIdAmts {
		recipesAndOccurrences = append(recipesAndOccurrences, RecipeAndOccurrence{ID: key, Occurrences: val})
	}

	sort.Sort(byOccurrences(recipesAndOccurrences))
	fmt.Println(recipesAndOccurrences)
	return myResponse, nil
}

type byOccurrences []RecipeAndOccurrence

type RecipeAndOccurrence struct {
	ID int
	Occurrences int
}

func (r byOccurrences) Len() int {
	return len(r)
}

func (r byOccurrences) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byOccurrences) Less(i, j int) bool {
	return r[i].Occurrences < r[j].Occurrences
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
// zip main.zip main
// aws lambda update-function-code --function-name mealmatesDriver --zip-file fileb://./main.zip
