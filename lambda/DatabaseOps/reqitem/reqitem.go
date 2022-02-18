package reqitem

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RequestItem interface {
	BuildWriteRequest() (map[string][]types.WriteRequest, error)
	BuildDeleteItem() (dynamodb.DeleteItemInput, error)
	BuildUpdateItem() (dynamodb.UpdateItemInput, error)
	BuildScanItem() (dynamodb.ScanInput, error)
	BuildGetItem() (dynamodb.GetItemInput, error)
}