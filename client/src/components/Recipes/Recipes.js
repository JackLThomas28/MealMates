import "./Recipe.css";
import RecipeList from "./RecipeList";

const Recipes = (props) => {
  return (
    <div className="recipes">
      <RecipeList items={props.items} />
    </div>
  );
};

export default Recipes;
