package ops

import (
	"context"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// Local
	"mealmates.com/lambda/DatabaseOps/reqitem"
)

func Scan(item reqitem.RequestItem) (*dynamodb.ScanOutput, error) {
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
	scanItem, err = item.BuildScanItem()

	// Only continue if there are no errors
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	// Scan with the item
	return svc.Scan(context.TODO(), &scanItem)
}