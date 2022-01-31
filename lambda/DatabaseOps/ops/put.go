package ops

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	// Local Packages
	"mealmates.com/lambda/DatabaseOps/structs"
)

func Put(recipe structs.Recipe) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

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

	out, err := svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			"ingredients": {
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
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

