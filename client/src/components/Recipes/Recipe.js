import "./Recipe.css";
import Ingredients from "./Ingredients/Ingredients";
import Image from "../UI/Image";

const Recipe = (props) => {
  return (
    <li className="recipe">
      <div>
        <h2>{props.name}</h2>
        <Image
          HEIGHT={200}
          src={props.imgSrc}
          height={props.imgHeight}
          width={props.imgWidth}
        />
        <h3>{props.servings}</h3>
        <Ingredients items={props.ingredients} />
      </div>
    </li>
  );
};

export default Recipe;
