package ops

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ingredient "mealmates.com/lambda/DatabaseOps/ingredient"
	"mealmates.com/lambda/DatabaseOps/recipe"
)

type DeleteItem struct {
	Recipe     recipe.Recipe
	Ingredient ingredient.Ingredient
}

func Delete(item DeleteItem, table string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	// Only continue if there are no errors
	if err != nil {
		return err
	}

	// Get the corresponding DeleteItem based on the recieved table name
	var deleteItem dynamodb.DeleteItemInput
	if table == recipe.TABLE_NAME {
		deleteItem, err = recipe.BuildDeleteItem(item.Recipe)
	} else if table == ingredient.TABLE_NAME {
		deleteItem, err = ingredient.BuildDeleteItem(item.Ingredient)
	}

	// Only continue if there are no errors
	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Delete the item
	_, err = svc.DeleteItem(context.TODO(), &deleteItem)
	return err
}
