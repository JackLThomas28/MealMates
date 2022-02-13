package ingredient

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

const TABLE_NAME = "Ingredients"

func BuildWriteRequest(ingredient Ingredient) (map[string][]types.WriteRequest, error) {
	writeRequests := make(map[string][]types.WriteRequest)
	writeRequests[TABLE_NAME] = make([]types.WriteRequest, 0)

	ingrRecipeIDs, err := attributevalue.MarshalList(ingredient.RecipeIDs)
	if err != nil {
		return map[string][]types.WriteRequest{}, err
	}

	return map[string][]types.WriteRequest{
		TABLE_NAME: {
			{
				PutRequest: &types.PutRequest{
					Item: map[string]types.AttributeValue{
						"name":      &types.AttributeValueMemberS{Value: ingredient.Name},
						"recipeIds": &types.AttributeValueMemberL{Value: ingrRecipeIDs},
					},
				},
			},
		},
	}, nil
}

func BuildDeleteItem(ingredient Ingredient) (dynamodb.DeleteItemInput, error) {
	return dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberN{Value: ingredient.Name},
		},
	}, nil
}

// Update the recipeIds list
func BuildUpdateItem(ingredient Ingredient) (dynamodb.UpdateItemInput, error) {
	// Only update the recipeIds if the recipeId is NOT in the list
	return dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: ingredient.Name},
		},
		UpdateExpression: aws.String("set recipeIds[0] = :recipeId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":recipeId": &types.AttributeValueMemberN{Value: strconv.Itoa(ingredient.RecipeID)},
		},
		ConditionExpression: aws.String("NOT contains(recipeIds, :recipeId)"),
	}, nil
}
