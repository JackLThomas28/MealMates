package recipe

import (
	"strconv"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Rating struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type Nutrition struct {
	Calories string `json:"calories"`
}

type Ingredient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
	Raw    string  `json:"raw"`
}

type Recipe struct {
	ID                int          `json:"id"`
	Name              string       `json:"name"`
	Image             Image        `json:"image"`
	Description       string       `json:"description"`
	PrepTime          string       `json:"prepTime"`
	CookTime          string       `json:"cookTime"`
	TotalTime         string       `json:"totalTime"`
	RecipeYield       string       `json:"recipeYield"`
	Ingredients       []string     `json:"ingredients"`
	Instructions      []string     `json:"instructions"`
	Categories        []string     `json:"categories"`
	Rating            Rating       `json:"rating"`
	Nutrition         Nutrition    `json:"nutrition"`
	ParsedIngredients []Ingredient `json:"parsedIngredients"`
}

const TABLE_NAME = "AllRecipes"

func BuildWriteRequest(recipe Recipe) (map[string][]types.WriteRequest, error) {
	recipeIngredients, err := attributevalue.MarshalList(recipe.Ingredients)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeInstructions, err := attributevalue.MarshalList(recipe.Instructions)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeCategories, err := attributevalue.MarshalList(recipe.Categories)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeRating, err := attributevalue.MarshalMap(recipe.Rating)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeImage, err := attributevalue.MarshalMap(recipe.Image)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeNutrition, err := attributevalue.MarshalMap(recipe.Nutrition)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}

	var parsedIngredientMaps []map[string]types.AttributeValue
	for _, ing := range recipe.ParsedIngredients {
		ingredientMap, err := attributevalue.MarshalMap(ing)
		if err != nil {
			return map[string][]types.WriteRequest{}, err
		} else {
			parsedIngredientMaps = append(parsedIngredientMaps, ingredientMap)
		}
	}
	parsedIngredients, err := attributevalue.MarshalList(parsedIngredientMaps)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}

	return map[string][]types.WriteRequest{
		TABLE_NAME: {
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
	}, nil
}

func BuildDeleteItem(recipe Recipe) (dynamodb.DeleteItemInput, error) {
	return dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(recipe.ID)},
		},
	}, nil
}

// Update the ID
func BuildUpdateItem(recipe Recipe) (dynamodb.UpdateItemInput, error) {
	return dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(recipe.ID)},
		},
		UpdateExpression: aws.String("set id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: strconv.Itoa(recipe.ID)},
		},
		// ConditionExpression: aws.String("id = :id"),
	}, nil
}
