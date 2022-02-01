import "./App.css";
import Header from "./components/Header";
import Search from "./components/Search/Search";
import Recipes from "./components/Recipes/Recipes";
import Footer from "./components/Footer";

/* Developement Only */
import recipes from "./data/recipes";
import { useState } from "react";
/* ***************** */

function App() {
  const [meals, setMeals] = useState(recipes);

  const searchHandler = (recipe) => {
    setMeals((prevRecipes) => {
      return [recipe, ...prevRecipes];
    });
  };

  return (
    <div>
      <Header />
      <Search onSearch={searchHandler} />
      <Recipes items={meals} />
      <Footer />
    </div>
  );
}

export default App;
