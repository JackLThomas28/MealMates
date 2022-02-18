package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// Local
	"mealmates.com/lambda/DatabaseOps/reqitem"
)

func Delete(item reqitem.RequestItem) error {
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
	deleteItem, err = item.BuildDeleteItem()

	// Only continue if there are no errors
	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Delete the item
	_, err = svc.DeleteItem(context.TODO(), &deleteItem)
	return err
}
