import Navbar from "../components/navbar";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import opggLogo from "../icons/hwei.png";
import './home.css';

export default function Home() {
    const SearchBar = () => {
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
            <form className="searchbar" onSubmit={searchSummoner}>
                <div className="region-selector">
                    <label htmlFor="regionTag" className="label" style={{width: "194px", height: "16px"}}>Region</label>
                    <select className="region-dropdown" defaultValue="na">
                        <option value="na">North America</option>
                        <option value="euw">Europe West</option>
                        <option value="kr">Korea</option>
                    </select>
                </div>
                <div className="search-wrapper">
                    <label htmlFor="searchHome" className="label">Search</label>
                    <input type="search" value={searchTerm} onChange={handleChange} placeholder="Game Name + #NA1" className="input-text" style={{width: "460px", height: "20px"}}/>
                </div>
                <button type="submit" className="gg-btn"></button>
            </form>
        );
    };
    return (
        <div>
            <Navbar />
            <div className="entire-page">
                <div className="logo-container">
                    <img src={opggLogo} height="224" alt="OP.GG" />
                </div>
                <SearchBar />
            </div>
        </div>
    )
}