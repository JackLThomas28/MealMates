import "./Search.css";
import SearchForm from "./SearchForm";

const Search = (props) => {
  return (
    <div>
      <SearchForm onSearch={props.onSearch} />
    </div>
  );
};

export default Search;
