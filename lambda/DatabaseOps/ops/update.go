package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	//Local
	"mealmates.com/lambda/DatabaseOps/reqitem"
)

func Update(item reqitem.RequestItem) error {
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
	updateItem, err = item.BuildUpdateItem()

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
