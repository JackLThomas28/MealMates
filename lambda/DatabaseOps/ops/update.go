package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	//Local
	"mealmates.com/lambda/DatabaseOps/ingredient"
	"mealmates.com/lambda/DatabaseOps/recipe"
)

type UpdateItem struct {
	Recipe     recipe.Recipe
	Ingredient ingredient.Ingredient
}

func Update(item UpdateItem, table string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	// Only continue if there are no errors
	if err != nil {
		return err
	}

	// Get the corresponding UpdateItem based on the recieved table name
	var updateItem dynamodb.UpdateItemInput
	if table == recipe.TABLE_NAME {
		updateItem, err = recipe.BuildUpdateItem(item.Recipe)
	} else if table == ingredient.TABLE_NAME {
		updateItem, err = ingredient.BuildUpdateItem(item.Ingredient)
	}

	// Only continue if there are no errors
	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Update the item
	_, err = svc.UpdateItem(context.TODO(), &updateItem)

	if err != nil {
		return err
	}

	return nil
}
