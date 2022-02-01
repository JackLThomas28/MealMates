import { useState } from "react";
import AWS from "aws-sdk";
import { Lambda } from "../../environment/Keys";

const SearchForm = (props) => {
  const [userInput, setUserInput] = useState({
    query: "",
  });

  const searchHandler = (event) => {
    // Make sure react is giving me the latest state w/ prevState
    setUserInput((prevState) => {
      return { ...prevState, query: event.target.value };
    });
  };

  const submitHandler = (event) => {
    // Prevent the page from reloading when the form is submitted
    event.preventDefault();

    const lambdaKeys = new Lambda();

    // Invoke the "getRecipe" AWS lambda function
    AWS.config.update({
      accessKeyId: lambdaKeys.getAccessKeyId(),
      secretAccessKey: lambdaKeys.getSecretAccessKey(),
      region: "us-east-1",
    });
    let lambda = new AWS.Lambda();
    let params = {
      FunctionName: "getRecipe",
      Payload: JSON.stringify({
        url: `${userInput.query}`,
      }),
    };
    lambda.invoke(params, (err, data) => {
      if (err) console.log("Error", err, err.stack);
      else {
        props.onSearch(JSON.parse(data.Payload).recipe);
      }
    });

    // Clear the search input box
    setUserInput((prevState) => {
      return { ...prevState, query: "" };
    });
  };

  return (
    <div className="searchform">
      <form onSubmit={submitHandler}>
        <input
          type="text"
          placeholder="Paste Recipe URL Here"
          onChange={searchHandler}
          value={userInput.query}
        />
        <button type="submit">Search</button>
      </form>
    </div>
  );
};

export default SearchForm;
