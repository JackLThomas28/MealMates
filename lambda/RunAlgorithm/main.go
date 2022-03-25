package main

import (
	"context"
	"sort"

	"github.com/JackLThomas28/MealMates/lambda/objects/ingredientdb"
	"github.com/JackLThomas28/MealMates/lambda/objects/recipeandfrequency"
	"github.com/aws/aws-lambda-go/lambda"
)

type RF recipeandfrequency.RecipeAndFrequency

type MyRequest struct {
	CntItemsToReturn int `json:"cntItemsToReturn"`
	Ingredients []ingredientdb.IngredientDB `json:"ingredients"`
}

type MyResponse struct {
	Recipes []RF `json:"recipes"`
}

func HandleRequest(ctx context.Context, request MyRequest) (MyResponse, error) {
	recipeIdAmts := make(map[int]int)
	for _, ingr := range request.Ingredients {
		for _, id := range ingr.RecipeIds {
			if recipeIdAmts[id] == 0 {
				recipeIdAmts[id] = 1
			} else {
				recipeIdAmts[id] += 1
			}
		}
	}

	var recipesAndFrequencies []RF
	for key, val := range recipeIdAmts {
		recipesAndFrequencies = append(recipesAndFrequencies, RF{ID: key, Frequency: val})
	}

	sort.Sort(byFrequency(recipesAndFrequencies))

	return MyResponse{Recipes: recipesAndFrequencies}, nil
}

type byFrequency []RF

// type RecipeAndOccurrence struct {
// 	ID int
// 	Occurrences int
// }

func (r byFrequency) Len() int {
	return len(r)
}

func (r byFrequency) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byFrequency) Less(i, j int) bool {
	return r[i].Frequency > r[j].Frequency
}

func main() {
	lambda.Start(HandleRequest)
}

// BUILD COMMAND:
// GOOS=linux go build main.go
// zip main.zip main
// aws lambda update-function-code --function-name runAlgorithm --zip-file fileb://./main.zip