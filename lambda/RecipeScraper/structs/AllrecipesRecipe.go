package structs

type ARItem struct {
	Id    string `json:"@id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ARListElement struct {
	Type     string `json:"@type"`
	Position string `json:"position"`
	Item     ARItem `json:"item"`
}

type ARBreadcrumbList struct {
	Context         string          `json:"@context"`
	Type            string          `json:"@type"`
	ItemListElement []ARListElement `json:"itemListElement"`
}

type ARImage struct {
	Type   string `json:"@type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ARStep struct {
	Type string `json:"@type"`
	Text string `json:"text"`
}

type ARPerson struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type ARNutrition struct {
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

type ARRating struct {
	Type         string  `json:"@type"`
	RatingValue  float32 `json:"ratingValue"`
	RatingCount  int     `json:"ratingCount"`
	ItemReviewed string  `json:"itemReviewed"`
	BestRating   string  `json:"bestRating"`
	WorstRating  string  `json:"worstRating"`
}

type ARReviewRating struct {
	Type        string `json:"@type"`
	WorstRating string `json:"worstRating"`
	BestRating  string `json:"bestRating"`
	RatingValue int    `json:"ratingValue"`
}

type ARAuthor struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	SameAs string `json:"sameAs"`
}

type ARReview struct {
	Type          string         `json:"@type"`
	DatePublished string         `json:"datePublished"`
	ReviewBody    string         `json:"reviewBody"`
	ReviewRating  ARReviewRating `json:"reviewRating"`
	Author        ARAuthor       `json:"author"`
}

type ARRecipe struct {
	ID                 int
	Context            string          `json:"@context"`
	Type               string          `json:"@type"`
	MainEntityOfPage   string          `json:"mainEntityOfPage"`
	Name               string          `json:"name"`
	Image              ARImage         `json:"image"`
	DatePublished      string          `json:"datePublished"`
	Description        string          `json:"description"`
	PrepTime           string          `json:"prepTime"`
	CookTime           string          `json:"cooktime"`
	TotalTime          string          `json:"totalTime"`
	RecipeYield        string          `json:"recipeYield"`
	RecipeIngredient   []string        `json:"recipeIngredient"`
	RecipeInstructions []ARStep        `json:"recipeInstructions"`
	RecipeCategory     []string        `json:"recipeCategory"`
	RecipeCuisine      []string        `json:"recipeCuisine"`
	Author             []ARPerson      `json:"author"`
	AggregateRating    ARRating        `json:"aggregateRating"`
	Nutrition          ARNutrition     `json:"nutrition"`
	Review             []ARReview      `json:"-"`
	ItemListElement    []ARListElement `json:"itemListElement"`
}
