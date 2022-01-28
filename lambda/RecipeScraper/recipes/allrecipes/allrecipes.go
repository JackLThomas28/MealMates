package allrecipes

import (
	"encoding/json"
	"errors"
	"mealmates.com/lambda/RecipeScraper/myutils"
	"regexp"
	"strconv"
	// "log"
)

type Item struct {
	Id    string `json:"@id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ListElement struct {
	Type     string `json:"@type"`
	Position string `json:"position"`
	Item     Item   `json:"item"`
}

type BreadcrumbList struct {
	Context         string        `json:"@context"`
	Type            string        `json:"@type"`
	ItemListElement []ListElement `json:"itemListElement"`
}

type Image struct {
	Type   string `json:"@type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Step struct {
	Type string `json:"@type"`
	Text string `json:"text"`
}

type Person struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type Nutrition struct {
	Type                  string `json:"@type"`
	Calories              string `json:"calories"`
	CarbohydrateContent   string `json:"carbohydrateContent"`
	CholesterolContent    string `json:"cholesterolContent"`
	FatContent            string `json:"fatContent"`
	FiberContent          string `json:"fiberContent"`
	ProteinContent        string `json:"proteinContent"`
	SaturatedFatContent   string `json:"saturatedFatContent"`
	ServingSize           int    `json:"servingSize"`
	SodiumContent         string `json:"sodiumContent"`
	SugarContent          string `json:"sugarContent"`
	TransFatContent       string `json:"transFatContent"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent"`
}

type Rating struct {
	Type         string  `json:"@type"`
	RatingValue  float32 `json:"ratingValue"`
	RatingCount  int     `json:"ratingCount"`
	ItemReviewed string  `json:"itemReviewed"`
	BestRating   string  `json:"bestRating"`
	WorstRating  string  `json:"worstRating"`
}

type ReviewRating struct {
	Type        string `json:"@type"`
	WorstRating string `json:"worstRating"`
	BestRating  string `json:"bestRating"`
	RatingValue int    `json:"ratingValue"`
}

type Author struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	SameAs string `json:"sameAs"`
}

type Review struct {
	Type          string       `json:"@type"`
	DatePublished string       `json:"datePublished"`
	ReviewBody    string       `json:"reviewBody"`
	ReviewRating  ReviewRating `json:"reviewRating"`
	Author        Author       `json:"author"`
}

type Recipe struct {
	ID                 int
	Context            string        `json:"@context"`
	Type               string        `json:"@type"`
	MainEntityOfPage   string        `json:"mainEntityOfPage"`
	Name               string        `json:"name"`
	Image              Image         `json:"image"`
	DatePublished      string        `json:"datePublished"`
	Description        string        `json:"description"`
	PrepTime           string        `json:"prepTime"`
	CookTime           string        `json:"cooktime"`
	TotalTime          string        `json:"totalTime"`
	RecipeYield        string        `json:"recipeYield"`
	RecipeIngredient   []string      `json:"recipeIngredient"`
	RecipeInstructions []Step        `json:"recipeInstructions"`
	RecipeCategory     []string      `json:"recipeCategory"`
	RecipeCuisine      []string      `json:"recipeCuisine"`
	Author             []Person      `json:"author"`
	AggregateRating    Rating        `json:"aggregateRating"`
	Nutrition          Nutrition     `json:"nutrition"`
	Review             []Review      `json:"-"`
	ItemListElement    []ListElement `json:"itemListElement"`
}

const URL = "https://www.allrecipes.com/recipe/"
const FILE_NAME = "allrecipes.json"

func GetRecipe(inputURL string) (Recipe, error) {
	// Get the recipe ID from the URL
	// Ex: "https://www.allrecipes.com/recipe/279984/air-fryer-sweet-and-spicy-roasted-carrots/"
	re, _ := regexp.Compile(URL)
	indices := re.FindStringIndex(inputURL)
	// Only save the last part of the url
	// Ex: "279984/air-fryer-sweet-and-spicy-roasted-carrots/"
	recipeId := inputURL[indices[1]:]
	re, _ = regexp.Compile("/")
	indices = re.FindStringIndex(recipeId)
	// Only save the numbers
	// Ex: "279984"
	recipeId = recipeId[:indices[0]]
	// Convert from string to int
	recipeIdInt, err := strconv.Atoi(recipeId)
	if err != nil {
		return Recipe{}, errors.New("Could not find/convert recipe ID")
	}

	node := myutils.GetHtmlNode(inputURL)

	const TYPE = "application/ld+json"
	n := myutils.GetElementByType(node, TYPE)
	if n == nil {
		return Recipe{}, errors.New("Could not get element by type.")
	}

	// Store the data in Recipe struct
	var recipe []Recipe
	json.Unmarshal([]byte(n.FirstChild.Data), &recipe)

	const RECIPEINDEX = 1
	if recipe[RECIPEINDEX].MainEntityOfPage == "" {
		return Recipe{}, errors.New("Could not get MainEntityOfPage property.")
	}

	// log.Printf("AllRecipes")
	// for i,_ := range recipe[RECIPEINDEX].RecipeIngredient {
	// 	log.Printf(recipe[RECIPEINDEX].RecipeIngredient[i])
	// }
	recipe[RECIPEINDEX].ID = recipeIdInt
	return recipe[RECIPEINDEX], nil
}
