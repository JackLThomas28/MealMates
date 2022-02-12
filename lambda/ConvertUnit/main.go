package main

import (
	"context"

	// Third Party
	lambda "github.com/aws/aws-lambda-go/lambda"

	convert "mealmates.com/lambda/ConvertUnit/Convert"
)

type MyEvent struct {
	Amount float64 `json:"amount"`
	To     string  `json:"to"`
	From   string  `json:"from"`
}

type MyResponse struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, request MyEvent) (MyResponse, error) {
	result, err := convert.Convert(request.Amount).From(request.From).To(request.To)
	return MyResponse{Result: result}, err
}

func main() {
	lambda.Start(HandleRequest)

	// result, err := converter.Convert(2).From("teaspoons").To("tablespoon")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result, err = converter.Convert(2).From("teaspoon").To("tablespoons")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result, err = converter.Convert(2).From("pound").To("tablespoon")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result, err = converter.Convert(20).From("pound").To("kilogram")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result, "kilograms")
}

// BUILD COMMAND:
// GOOS=linux go build main.go
