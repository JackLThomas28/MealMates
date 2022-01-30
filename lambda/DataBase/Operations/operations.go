package operations

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Ingredient struct {
	Name string `json:"name"`
	Amount float64 `json:"amount"`
	Unit string `json:"unit"`
	Raw string `json:"raw"`
}

type ImageJSON struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type RatingJSON struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type NutritionJSON struct {
	Calories string `json:"calories"`
}

type RecipeJSON struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Image        ImageJSON     `json:"image"`
	Description  string        `json:"description"`
	PrepTime     string        `json:"prepTime"`
	CookTime     string        `json:"cookTime"`
	TotalTime    string        `json:"totalTime"`
	RecipeYield  string        `json:"recipeYield"`
	Ingredients  []string      `json:"ingredients"`
	Instructions []string      `json:"instructions"`
	Categories   []string      `json:"categories"`
	Rating       RatingJSON    `json:"rating"`
	Nutrition    NutritionJSON `json:"nutrition"`
	ParsedIngredients []Ingredient `json:"parsedIngredients"`
}

func Put(recipe RecipeJSON) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
        o.Region = "us-east-1"
        return nil
    })
    if err != nil {
        panic(err)
    }

    svc := dynamodb.NewFromConfig(cfg)

	recipeIngredients, err := attributevalue.MarshalList(recipe.Ingredients)
	recipeInstructions, err := attributevalue.MarshalList(recipe.Instructions)
	recipeCategories, err := attributevalue.MarshalList(recipe.Categories)
	recipeRating, err := attributevalue.MarshalMap(recipe.Rating)
	recipeImage, err := attributevalue.MarshalMap(recipe.Image)
	recipeNutrition, err := attributevalue.MarshalMap(recipe.Nutrition)

	
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

    out, err := svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
        RequestItems: map[string][]types.WriteRequest{
            "ingredients": {
                {
                    PutRequest: &types.PutRequest{
                        Item: map[string]types.AttributeValue{
                            "id": &types.AttributeValueMemberN{Value: strconv.Itoa(recipe.ID)},
                            "name": &types.AttributeValueMemberS{Value: recipe.Name},
                            "image": &types.AttributeValueMemberM{Value: recipeImage},
                            "description": &types.AttributeValueMemberS{Value: recipe.Description},
                            "prepTime": &types.AttributeValueMemberS{Value: recipe.PrepTime},
                            "cookTime": &types.AttributeValueMemberS{Value: recipe.CookTime},
                            "totalTime": &types.AttributeValueMemberS{Value: recipe.TotalTime},
                            "recipeYield": &types.AttributeValueMemberS{Value: recipe.RecipeYield},
                            "ingredients": &types.AttributeValueMemberL{Value: recipeIngredients},
                            "instructions": &types.AttributeValueMemberL{Value: recipeInstructions},
                            "categories":  &types.AttributeValueMemberL{Value: recipeCategories},
                            "rating":  &types.AttributeValueMemberM{Value: recipeRating},
                            "nutrition":  &types.AttributeValueMemberM{Value: recipeNutrition},
                            "parsedIngredients":  &types.AttributeValueMemberL{Value: parsedIngredients},
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

