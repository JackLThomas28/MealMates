package reqitem

import (
	"errors"
	"fmt"
	"strconv"

	// Third Party
	"github.com/JackLThomas28/MealMates/lambda/objects/ingredientdb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const INGR_TABLE_NAME = "Ingredients"

type MyIngredientDB ingredientdb.IngredientDB
func (i MyIngredientDB) BuildWriteRequest() (map[string][]types.WriteRequest, error) {
	writeRequests := make(map[string][]types.WriteRequest)
	writeRequests[INGR_TABLE_NAME] = make([]types.WriteRequest, 0)

	ingrRecipeIDs, err := attributevalue.MarshalList(i.RecipeIds)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}

	return map[string][]types.WriteRequest{
		INGR_TABLE_NAME: {
			{
				PutRequest: &types.PutRequest{
					Item: map[string]types.AttributeValue{
						"name":      &types.AttributeValueMemberS{Value: i.Name},
						"recipeIds": &types.AttributeValueMemberL{Value: ingrRecipeIDs},
					},
				},
			},
		},
	}, nil
}

func (i MyIngredientDB) BuildDeleteItem() (dynamodb.DeleteItemInput, error) {
	return dynamodb.DeleteItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberN{Value: i.Name},
		},
	}, nil
}

// Update the recipeIds list
func (i MyIngredientDB) BuildUpdateItem() (dynamodb.UpdateItemInput, error) {
	ingrRecipeIDs, err := attributevalue.MarshalList(i.RecipeIds)
	if err != nil {
		return dynamodb.UpdateItemInput{}, err
	}
	// Only update the recipeIds if the recipeId is NOT in the list
	return dynamodb.UpdateItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: i.Name},
		},
		UpdateExpression: aws.String("set recipeIds = list_append(recipeIds, :recipeIds)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":recipeIds": &types.AttributeValueMemberL{Value: ingrRecipeIDs},
			":recipeId": &types.AttributeValueMemberN{Value: strconv.Itoa(i.RecipeIds[0])},
		},
		ConditionExpression: aws.String("NOT contains(recipeIds, :recipeId)"),
	}, nil
}

// Scan based on name
func (i MyIngredientDB) BuildScanItem() (dynamodb.ScanInput, error) {
	return dynamodb.ScanInput{
		TableName: aws.String(INGR_TABLE_NAME),
		FilterExpression: aws.String("contains(#name, :name)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: i.Name},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": "name",
		},
	}, nil
}

// Get based on name (primary key)
func (i MyIngredientDB) BuildGetItem() (dynamodb.GetItemInput, error) {
	return dynamodb.GetItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: i.Name},
		},
	}, nil
}

func extractIngredient(i *MyIngredientDB, result map[string]types.AttributeValue) error {
	name := result["name"].(*types.AttributeValueMemberS)
	err := attributevalue.Unmarshal(name, &i.Name)
	if err != nil {
		return err
	}

	ids := result["recipeIds"].(*types.AttributeValueMemberL)
	err = attributevalue.UnmarshalList(ids.Value, &i.RecipeIds)
	if err != nil {
		return err
	}

	return nil
}

func (i *MyIngredientDB) ParseResult(result map[string]types.AttributeValue) error {
	if result == nil {
		return errors.New("ingredient parseresult: could not locate ingredient")
	}
	err := extractIngredient(i, result)
	return err
}

func (i MyIngredientDB) ParseScanResults(results []map[string]types.AttributeValue) ([]RequestItem, error) {
	var ingrList []RequestItem
	var err error

	fmt.Println(results)

	for _, ingr := range results {
		curIngr := &MyIngredientDB{}
		err = extractIngredient(curIngr, ingr)
		ingrList = append(ingrList, curIngr)
	}
	return ingrList, err
}