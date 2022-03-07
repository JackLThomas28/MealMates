package mylambda

import (
	"context"
	"fmt"
	"os"

	// Third Party
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type staticCredentialsProvider struct {
	Value aws.Credentials
}

func (s staticCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	v := s.Value
	if v.AccessKeyID == "" || v.SecretAccessKey == "" {
		return aws.Credentials{
			Source: "Source Name",
		}, fmt.Errorf("static credentials are empty")
	}

	if len(v.Source) == 0 {
		v.Source = "Source Name"
	}

	return v, nil
}

type SymmetricCredentialAdaptor struct {
	SymmetricProvider aws.CredentialsProvider
}

func InvokeLambda(ctx context.Context, funcName string, reqPayload []byte) ([]byte, error) {
	credProvider := &staticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     os.Getenv("ACCESS_KEY"),
			SecretAccessKey: os.Getenv("SECRET_KEY"),
		},
	}
	
	options := lambda.Options{
		Region: "us-east-1",
		Credentials: credProvider,
	}
	client := lambda.New(options)

	params := lambda.InvokeInput{
		FunctionName: aws.String(funcName),
		Payload: reqPayload,
	}

	response, err := client.Invoke(ctx, &params)
	return response.Payload, err
}
