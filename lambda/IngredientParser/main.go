package main

import (
	"context"
	// "github.com/aws/aws-lambda-go/lambda"
	"fmt"

	// Local Packages
	"mealmates.com/lambda/IngredientParser/Parser"
)

type MyEvent struct {
	Ingredients []string `json:"ingredients"` 
}

func HandleRequest(ctx context.Context, request MyEvent) (string, error) {
	parser.ParseIngredients(request.Ingredients)
	return "", nil
}

func main() {
	/* Development Only */
	var ingredientList []string
	ingredientList = append(ingredientList, "8 ounces elbow macaroni")
	ingredientList = append(ingredientList, "1 cup sour cream")
	ingredientList = append(ingredientList, "1 cup cottage cheese")
	ingredientList = append(ingredientList, "1 dash Worcestershire sauce")
	ingredientList = append(ingredientList, "3 tablespoons minced onion")
	ingredientList = append(ingredientList, "1 teaspoon minced garlic")
	ingredientList = append(ingredientList, "½ cup seasoned dry bread crumbs")
	ingredientList = append(ingredientList, "cooking spray")
	ingredientList = append(ingredientList, "6 tablespoons hot pepper sauce")
	ingredientList = append(ingredientList, "1 ⅓ cup olive oil")
	ingredientList = append(ingredientList, "2 tablespoons garlic powder")
	ingredientList = append(ingredientList, "1 tablespoon freshly ground black pepper")
	ingredientList = append(ingredientList, "1 tablespoon paprika")
	ingredientList = append(ingredientList, "1 ½ teaspoons salt")
	ingredientList = append(ingredientList, "8 potatoes, cut into 1/2-inch cubes")
	ingredientList = append(ingredientList, "2 pounds skinless, boneless chicken breast halves, cut into 1/2-inch cubes")
	ingredientList = append(ingredientList, "2 cups shredded Mexican cheese blend (such as Great Value Fiesta Blend®)")
	ingredientList = append(ingredientList, "1 cup crumbled cooked bacon")
	ingredientList = append(ingredientList, "1 cup diced green onions")
	ingredientList = append(ingredientList, "1 (3 pound) whole chicken")
	ingredientList = append(ingredientList, "1 (14 ounce) jar butter beans, rinsed and drained")
	ingredientList = append(ingredientList, "3 (10.75 ounce) cans  condensed cream of mushroom soup")
	ingredientList = append(ingredientList, "10.75 (3 ounce) cans  condensed cream of mushroom soup")

	var event MyEvent
	event.Ingredients = ingredientList

	_, err := HandleRequest(nil, event)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("- - - - - - -")
	}
	/* **************** */

	// lambda.Start(HandleRequest)
}