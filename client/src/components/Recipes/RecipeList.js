import "./RecipeList.css";
import Recipe from "./Recipe";

const RecipeList = (props) => {
  if (props.items.length === 0) {
    return (
      <h3>Enter a recipe URL in the search box above to see similar recipes</h3>
    );
  }
  return (
    <ul className="recipe-list">
      {props.items.map((recipe) => (
        <Recipe
          key={recipe.id}
          name={recipe.name}
          imgSrc={recipe.image.url}
          imgHeight={recipe.image.height}
          imgWidth={recipe.image.width}
          description={recipe.description}
          servings={recipe.recipeYield}
          ingredients={recipe.ingredients}
        />
      ))}
    </ul>
  );
};

export default RecipeList;
