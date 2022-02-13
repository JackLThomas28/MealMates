package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	// Local
	ingredient "mealmates.com/lambda/DatabaseOps/ingredient"
	recipe "mealmates.com/lambda/DatabaseOps/recipe"
)

type PutItem struct {
	Recipe     recipe.Recipe
	Ingredient ingredient.Ingredient
}

func Put(item PutItem, table string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	// Only continue if there are no errors
	if err != nil {
		return err
	}

	// Get the corresponding WriteRequest based on the received table name
	var writeRequest map[string][]types.WriteRequest
	if table == recipe.TABLE_NAME {
		writeRequest, err = recipe.BuildWriteRequest(item.Recipe)
	} else if table == ingredient.TABLE_NAME {
		writeRequest, err = ingredient.BuildWriteRequest(item.Ingredient)
	}

	// Only continue if there are no errors
	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Write the item
	_, err = svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: writeRequest,
	})
	return err
}
