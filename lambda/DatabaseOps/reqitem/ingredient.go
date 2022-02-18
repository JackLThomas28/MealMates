package reqitem

import (
	"strconv"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Ingredient struct {
	Name      string `json:"name"`
	RecipeIDs []int  `json:"recipeIds"`
	RecipeID  int    `json:"recipeId"`
}

const INGR_TABLE_NAME = "Ingredients"

func (i Ingredient) BuildWriteRequest() (map[string][]types.WriteRequest, error) {
	writeRequests := make(map[string][]types.WriteRequest)
	writeRequests[INGR_TABLE_NAME] = make([]types.WriteRequest, 0)

	ingrRecipeIDs, err := attributevalue.MarshalList(i.RecipeIDs)
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

func (i Ingredient) BuildDeleteItem() (dynamodb.DeleteItemInput, error) {
	return dynamodb.DeleteItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberN{Value: i.Name},
		},
	}, nil
}

// Update the recipeIds list
func (i Ingredient) BuildUpdateItem() (dynamodb.UpdateItemInput, error) {
	// Only update the recipeIds if the recipeId is NOT in the list
	return dynamodb.UpdateItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: i.Name},
		},
		UpdateExpression: aws.String("set recipeIds[0] = :recipeId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":recipeId": &types.AttributeValueMemberN{Value: strconv.Itoa(i.RecipeID)},
		},
		ConditionExpression: aws.String("NOT contains(recipeIds, :recipeId)"),
	}, nil
}

// Scan based on name
func (i Ingredient) BuildScanItem() (dynamodb.ScanInput, error) {
	return dynamodb.ScanInput{
		TableName: aws.String(INGR_TABLE_NAME),
		FilterExpression: aws.String("contains(name, :name)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: i.Name},
		},
	}, nil
}

// Get based on name (primary key)
func (i Ingredient) BuildGetItem() (dynamodb.GetItemInput, error) {
	return dynamodb.GetItemInput{
		TableName: aws.String(INGR_TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: i.Name},
		},
	}, nil
}