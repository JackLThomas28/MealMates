package reqitem

import (
	"fmt"
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
	for _, ingr := range r.ParsedIngredients {
		// thisIngr, err := attributevalue.MarshalMap(ingr)
		if err != nil {
			return map[string][]types.WriteRequest{}, err
		}
		// parsedIngredientMaps = append(parsedIngredientMaps, thisIngr)
		parsedIngredientMaps = append(parsedIngredientMaps, map[string]types.AttributeValue{
			"name":   &types.AttributeValueMemberS{Value: ingr.Name},
			"amount": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", ingr.Amount)},
			"unit":   &types.AttributeValueMemberS{Value: ingr.Unit},
			"raw":    &types.AttributeValueMemberS{Value: ingr.Raw},
		})
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
		TableName:        aws.String(RECIPE_TABLE_NAME),
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

func (r *Recipe) ParseResult(result map[string]types.AttributeValue) error {
	ingredients := result["ingredients"].(*types.AttributeValueMemberL)
	err := attributevalue.UnmarshalList(ingredients.Value, &r.Ingredients)
	if err != nil {
		return err
	}
	instructions := result["instructions"].(*types.AttributeValueMemberL)
	err = attributevalue.UnmarshalList(instructions.Value, &r.Instructions)
	if err != nil {
		return err
	}
	categories := result["categories"].(*types.AttributeValueMemberL)
	err = attributevalue.UnmarshalList(categories.Value, &r.Categories)
	if err != nil {
		return err
	}
	rating := result["rating"].(*types.AttributeValueMemberM)
	err = attributevalue.UnmarshalMap(rating.Value, &r.Rating)
	if err != nil {
		return err
	}
	image := result["image"].(*types.AttributeValueMemberM)
	err = attributevalue.UnmarshalMap(image.Value, &r.Image)
	if err != nil {
		return err
	}
	nutrition := result["nutrition"].(*types.AttributeValueMemberM)
	err = attributevalue.UnmarshalMap(nutrition.Value, &r.Nutrition)
	if err != nil {
		return err
	}
	id := result["id"].(*types.AttributeValueMemberN)
	var idStr string
	err = attributevalue.Unmarshal(id, &idStr)
	if err != nil {
		return err
	}
	r.ID, err = strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	name := result["name"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(name, &r.Name)
	if err != nil {
		return err
	}
	description := result["description"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(description, &r.Description)
	if err != nil {
		return err
	}
	prepTime := result["prepTime"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(prepTime, &r.PrepTime)
	if err != nil {
		return err
	}
	cookTime := result["cookTime"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(cookTime, &r.CookTime)
	if err != nil {
		return err
	}
	totalTime := result["totalTime"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(totalTime, &r.TotalTime)
	if err != nil {
		return err
	}
	recipeYield := result["recipeYield"].(*types.AttributeValueMemberS)
	err = attributevalue.Unmarshal(recipeYield, &r.RecipeYield)
	if err != nil {
		return err
	}

	// Unmarshal list
	prsedIngrdnts := result["parsedIngredients"].(*types.AttributeValueMemberL)
	if err != nil {
		return err
	}
	
	parsedIngredients := prsedIngrdnts.Value

	for _, ingr := range parsedIngredients {
		var curIngr ingredient
		ingrMap := ingr.(*types.AttributeValueMemberM)

		// Ingredient Name
		ingrName := ingrMap.Value["name"].(*types.AttributeValueMemberM)
		err = attributevalue.Unmarshal(ingrName.Value["Value"], &curIngr.Name)
		if err != nil {
			return err
		}

		// Ingredient Unit
		ingrUnit := ingrMap.Value["unit"].(*types.AttributeValueMemberM)
		err = attributevalue.Unmarshal(ingrUnit.Value["Value"], &curIngr.Unit)
		if err != nil {
			return err
		}

		// Ingredient Raw
		ingrRaw := ingrMap.Value["raw"].(*types.AttributeValueMemberM)
		err = attributevalue.Unmarshal(ingrRaw.Value["Value"], &curIngr.Raw)
		if err != nil {
			return err
		}

		// Ingredient Amount
		ingrAmount := ingrMap.Value["amount"].(*types.AttributeValueMemberM)
		var amtStr string
		err = attributevalue.Unmarshal(ingrAmount.Value["Value"], &amtStr)
		if err != nil {
			return err
		}
		curIngr.Amount, err = strconv.ParseFloat(amtStr, 64)
		if err != nil {
			return err
		}

		r.ParsedIngredients = append(r.ParsedIngredients, curIngr)
	}
	return nil
}
