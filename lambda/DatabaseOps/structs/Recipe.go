package structs

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Rating struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type Nutrition struct {
	Calories string `json:"calories"`
}

type Recipe struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Image             Image         `json:"image"`
	Description       string        `json:"description"`
	PrepTime          string        `json:"prepTime"`
	CookTime          string        `json:"cookTime"`
	TotalTime         string        `json:"totalTime"`
	RecipeYield       string        `json:"recipeYield"`
	Ingredients       []string      `json:"ingredients"`
	Instructions      []string      `json:"instructions"`
	Categories        []string      `json:"categories"`
	Rating            Rating        `json:"rating"`
	Nutrition         Nutrition     `json:"nutrition"`
	ParsedIngredients []Ingredient  `json:"parsedIngredients"`
}
