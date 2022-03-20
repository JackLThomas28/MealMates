package ingredient

type IngredientDB struct {
	Name string `json:"name"`
	RecipeIds []int `json:"recipeIds"`
	// TODO: Remove RecipeId
	RecipeId int `json:"recipeId"`
}