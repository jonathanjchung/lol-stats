import { useNavigate } from "react-router-dom";
import { useState } from "react";


/*
* This search will be for the other pages, not the home page
*/
const Search = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const navigate = useNavigate();

    const handleChange = (event) => {
        setSearchTerm(event.target.value);
    }

    const searchSummoner = (event) => {
        event.preventDefault();
        if (searchTerm.trim() != "") {
            navigate(`summoners/${searchTerm}/NA1`);
        }
    }
    return (
        <form onSubmit={searchSummoner}>
            <input value={searchTerm} onChange={handleChange} placeholder="Game Name + #NA1" />
            <button type="submit">Search</button>
        </form>
    );
};

export default Search;