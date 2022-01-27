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
          key={recipe.ID}
          name={recipe.Name}
          imgSrc={recipe.Image.URL}
          imgHeight={recipe.Image.Height}
          imgWidth={recipe.Image.Width}
          description={recipe.Description}
          servings={recipe.RecipeYield}
          ingredients={recipe.Ingredients}
        />
      ))}
    </ul>
  );
};

export default RecipeList;
