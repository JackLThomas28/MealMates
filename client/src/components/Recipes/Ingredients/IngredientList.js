import "./IngredientList.css";
import Ingredient from "./Ingredient";

const IngredientList = (props) => {
  return (
    <ul>
      <div>
        {props.items.map((ingredient, index) => (
          <Ingredient key={index} description={ingredient} />
        ))}
      </div>
    </ul>
  );
};

export default IngredientList;
