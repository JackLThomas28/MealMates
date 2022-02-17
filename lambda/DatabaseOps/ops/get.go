package ops

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"mealmates.com/lambda/DatabaseOps/ingredient"
	"mealmates.com/lambda/DatabaseOps/recipe"
)

type GetItem struct {
	Recipe recipe.Recipe
	Ingredient ingredient.Ingredient
}

func Get(item UpdateItem, table string) (*dynamodb.GetItemOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	// Get the corresponding GetItem based on the recieved table name
	var getItem dynamodb.GetItemInput
	if table == recipe.TABLE_NAME {
		getItem, err = recipe.BuildGetItem(item.Recipe)
	} else if table == ingredient.TABLE_NAME {
		getItem, err = ingredient.BuildGetItem(item.Ingredient)
	}

	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Get the item
	return svc.GetItem(context.TODO(), &getItem)
}