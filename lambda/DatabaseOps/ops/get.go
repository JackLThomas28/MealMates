package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// Local
	"mealmates.com/lambda/DatabaseOps/reqitem"
)

func Get(item reqitem.RequestItem) (*dynamodb.GetItemOutput, error) {
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
	getItem, err = item.BuildGetItem()

	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Get the item
	return svc.GetItem(context.TODO(), &getItem)
}