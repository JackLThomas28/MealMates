package ingredient

type Ingredient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
	Raw    string  `json:"raw"`
}

type DBIngredient struct {
	Name string `json:"name"`
	RecipeIds []int `json:"recipeIds"`
	// TODO: Remove RecipeId
	RecipeId int `json:"recipeId"`
}