package reqitem

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

type ingredient struct {
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
	ParsedIngredients []ingredient `json:"parsedIngredients"`
}

const RECIPE_TABLE_NAME = "AllRecipes"

func (r Recipe) BuildWriteRequest() (map[string][]types.WriteRequest, error) {
	recipeIngredients, err := attributevalue.MarshalList(r.Ingredients)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeInstructions, err := attributevalue.MarshalList(r.Instructions)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeCategories, err := attributevalue.MarshalList(r.Categories)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeRating, err := attributevalue.MarshalMap(r.Rating)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeImage, err := attributevalue.MarshalMap(r.Image)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}
	recipeNutrition, err := attributevalue.MarshalMap(r.Nutrition)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}

	var parsedIngredientMaps []map[string]types.AttributeValue
	for _, ing := range r.ParsedIngredients {
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
		RECIPE_TABLE_NAME: {
			{
				PutRequest: &types.PutRequest{
					Item: map[string]types.AttributeValue{
						"id":                &types.AttributeValueMemberN{Value: strconv.Itoa(r.ID)},
						"name":              &types.AttributeValueMemberS{Value: r.Name},
						"image":             &types.AttributeValueMemberM{Value: recipeImage},
						"description":       &types.AttributeValueMemberS{Value: r.Description},
						"prepTime":          &types.AttributeValueMemberS{Value: r.PrepTime},
						"cookTime":          &types.AttributeValueMemberS{Value: r.CookTime},
						"totalTime":         &types.AttributeValueMemberS{Value: r.TotalTime},
						"recipeYield":       &types.AttributeValueMemberS{Value: r.RecipeYield},
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

func (r Recipe) BuildDeleteItem() (dynamodb.DeleteItemInput, error) {
	return dynamodb.DeleteItemInput{
		TableName: aws.String(RECIPE_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(r.ID)},
		},
	}, nil
}

// Update the ID
func (r Recipe) BuildUpdateItem() (dynamodb.UpdateItemInput, error) {
	return dynamodb.UpdateItemInput{
		TableName: aws.String(RECIPE_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(r.ID)},
		},
		UpdateExpression: aws.String("set id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberN{Value: strconv.Itoa(r.ID)},
		},
		// ConditionExpression: aws.String("id = :id"),
	}, nil
}

// Scan based on name
func (r Recipe) BuildScanItem() (dynamodb.ScanInput, error) {
	return dynamodb.ScanInput{
		TableName: aws.String(RECIPE_TABLE_NAME),
		FilterExpression: aws.String("contains(name, :name)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: r.Name},
		},
	}, nil
}

// Get based on id (primary key)
func (r Recipe) BuildGetItem() (dynamodb.GetItemInput, error) {
	return dynamodb.GetItemInput{
		TableName: aws.String(RECIPE_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(r.ID)},
		},
	}, nil
}