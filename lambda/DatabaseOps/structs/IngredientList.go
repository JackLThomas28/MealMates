package structs

type MyIngredient struct {
	Name string `json:"name"`
	RecipeIDs []int `json:"recipeIDs"`
}