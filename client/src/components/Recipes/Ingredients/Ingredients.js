import "./Ingredients.css";
import IngredientList from "./IngredientList";

const Ingredients = (props) => {
  return (
    <div className="ingredients">
      <IngredientList items={props.items} />
    </div>
  );
};

export default Ingredients;
