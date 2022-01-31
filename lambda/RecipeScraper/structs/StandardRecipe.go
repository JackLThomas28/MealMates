package structs

type StandardImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type StandardRating struct {
	RatingValue float32 `json:"ratingValue"`
	RatingCount int     `json:"ratingCount"`
}

type StandardNutrition struct {
	Calories string `json:"calories"`
}

type StandardRecipe struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Image        StandardImage     `json:"image"`
	Description  string    `json:"description"`
	PrepTime     string    `json:"prepTime"`
	CookTime     string    `json:"cookTime"`
	TotalTime    string    `json:"totalTime"`
	RecipeYield  string    `json:"recipeYield"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Categories   []string  `json:"categories"`
	Rating       StandardRating    `json:"rating"`
	Nutrition    StandardNutrition `json:"nutrition"`
}
