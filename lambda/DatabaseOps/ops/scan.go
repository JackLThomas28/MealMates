package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// Local
	ingredient "mealmates.com/lambda/DatabaseOps/ingredient"
	recipe "mealmates.com/lambda/DatabaseOps/recipe"
)

type ScanItem struct {
	Recipe recipe.Recipe
	Ingredient ingredient.Ingredient
}

func Scan(item ScanItem, table string) (*dynamodb.ScanOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	// Get the corresponding ScanItem based on the recieved table name
	var scanItem dynamodb.ScanInput
	if table == recipe.TABLE_NAME {
		scanItem, err = recipe.BuildScanItem(item.Recipe)
	} else if table == ingredient.TABLE_NAME {
		scanItem, err = ingredient.BuildScanItem(item.Ingredient)
	}

	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Scan with the item
	return svc.Scan(context.TODO(), &scanItem)
}