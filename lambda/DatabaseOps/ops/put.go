package ops

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	// Local Packages
	"mealmates.com/lambda/DatabaseOps/structs"
)

const INGREDIENT_TABLE_NAME = "Ingredients"
const ALLRECIPES_TABLE_NAME = "AllRecipes"

type PutItem struct {
	Recipe structs.Recipe
	Ingredients []structs.MyIngredient
}

func buildRecipesPutRequestItem(recipe structs.Recipe) map[string][]types.WriteRequest {
	recipeIngredients, err := attributevalue.MarshalList(recipe.Ingredients)
	if err != nil {
		fmt.Println(err)
	}
	recipeInstructions, err := attributevalue.MarshalList(recipe.Instructions)
	if err != nil {
		fmt.Println(err)
	}
	recipeCategories, err := attributevalue.MarshalList(recipe.Categories)
	if err != nil {
		fmt.Println(err)
	}
	recipeRating, err := attributevalue.MarshalMap(recipe.Rating)
	if err != nil {
		fmt.Println(err)
	}
	recipeImage, err := attributevalue.MarshalMap(recipe.Image)
	if err != nil {
		fmt.Println(err)
	}
	recipeNutrition, err := attributevalue.MarshalMap(recipe.Nutrition)
	if err != nil {
		fmt.Println(err)
	}

	var parsedIngredientMaps []map[string]types.AttributeValue
	for _, ing := range recipe.ParsedIngredients {
		ingredientMap, err := attributevalue.MarshalMap(ing)
		if err != nil {
			fmt.Println(err)
		} else {
			parsedIngredientMaps = append(parsedIngredientMaps, ingredientMap)
		}
	}
	parsedIngredients, err := attributevalue.MarshalList(parsedIngredientMaps)
	if err != nil {
		fmt.Println(err)
	}

	return map[string][]types.WriteRequest{
		ALLRECIPES_TABLE_NAME: {
			{
				PutRequest: &types.PutRequest{
					Item: map[string]types.AttributeValue{
						"id":                &types.AttributeValueMemberN{Value: strconv.Itoa(recipe.ID)},
						"name":              &types.AttributeValueMemberS{Value: recipe.Name},
						"image":             &types.AttributeValueMemberM{Value: recipeImage},
						"description":       &types.AttributeValueMemberS{Value: recipe.Description},
						"prepTime":          &types.AttributeValueMemberS{Value: recipe.PrepTime},
						"cookTime":          &types.AttributeValueMemberS{Value: recipe.CookTime},
						"totalTime":         &types.AttributeValueMemberS{Value: recipe.TotalTime},
						"recipeYield":       &types.AttributeValueMemberS{Value: recipe.RecipeYield},
						"ingredients":       &types.AttributeValueMemberL{Value: recipeIngredients},
						"instructions":      &types.AttributeValueMemberL{Value: recipeInstructions},
						"categories":        &types.AttributeValueMemberL{Value: recipeCategories},
						"rating":            &types.AttributeValueMemberM{Value: recipeRating},
						"nutrition":         &types.AttributeValueMemberM{Value: recipeNutrition},
						"parsedIngredients": &types.AttributeValueMemberL{Value: parsedIngredients},
					},
				},
			},
		},
	}
}

func buildIngredientListRequestItem(ingredients []structs.MyIngredient) map[string][]types.WriteRequest {
	writeRequests := make(map[string][]types.WriteRequest)
	writeRequests[INGREDIENT_TABLE_NAME] = make([]types.WriteRequest, 0)
	
	for i, ingr := range ingredients {
		if i >= 25 {
			break
		}

		recipeIDs, err := attributevalue.MarshalList(ingr.RecipeIDs)
		if err != nil {
			fmt.Println(err)
		}

		wr := types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: map[string]types.AttributeValue{
					"name": &types.AttributeValueMemberS{Value: ingr.Name},
					"recipeIDs": &types.AttributeValueMemberL{Value: recipeIDs},
				},
			},
		}
		writeRequests[INGREDIENT_TABLE_NAME] = append(writeRequests[INGREDIENT_TABLE_NAME], wr)
	}
	return writeRequests
}

func Put(item PutItem, table string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	var writeRequest map[string][]types.WriteRequest
	if table == ALLRECIPES_TABLE_NAME {
		writeRequest = buildRecipesPutRequestItem(item.Recipe)
	} else if table == INGREDIENT_TABLE_NAME {
		writeRequest = buildIngredientListRequestItem(item.Ingredients)
	}
	
	_, err = svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: writeRequest,
	})
	if err != nil {
		panic(err)
	}
}

