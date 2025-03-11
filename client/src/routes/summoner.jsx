import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import Navbar from "../components/navbar";
import MatchData from "../components/matchdata";
import profileIcon555 from "../icons/profileicon555.png";
import './summoner.css'

export default function Summoner() {
    const { summonerName, tagLine } = useParams();
    const [accountInfo, setAccountInfo] = useState(null);
    const [summonerInfo, setSummonerInfo] = useState(null);
    const [matchHistory, setMatchHistory] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const url = `http://localhost:8080/summoners?summonerName=${summonerName}&tag=${tagLine}`;

    const renderMatches = () => {
        if(!matchHistory || matchHistory.length === 0) {
            return <p>No matches found</p>;
        }
        else {
            return (
                <div className="match-history">
                    {matchHistory.map((m) => (
                        <div className="match-wrapper">
                            <MatchData key={m.info.gameId} match={m} summName={accountInfo.gameName} />
                        </div>
                    ))}
                </div>
            )
        }
    }

    useEffect(() => {
        console.log(`Looking at summoner: ${summonerName}#${tagLine}`); //
        setLoading(true);
        fetch(url)
        .then((res) => {
            if (!res.ok) {
                throw new Error('http error: ' + res.status);
            }
            return res.json();
        })
        .then((data) => {
            console.log(data)
            setAccountInfo(data.Account);
            setSummonerInfo(data.Summoner);
            setMatchHistory(data.Matches);
            setLoading(false);
        })
        .catch((err) => {
            console.error(err);
            setError(err);
            setLoading(false);
        });
    }, [url]);

    return (
        <div style={{ backgroundColor: "white" }}>
            <Navbar />
            {loading && <div>Loading...</div>}
            {error && <div>Error: {error.message}</div>}
            
            {!loading && !error && (
                <>
                    <div className="content-header">
                        <div className="wrapper">
                            <div className="profile-info">
                                <div className="summoner-icon">
                                    <img src={profileIcon555} alt="Summoner icon"/>
                                    <div className="level">
                                        <span className="level">
                                            {summonerInfo.summonerLevel}
                                        </span>
                                    </div>
                                </div>
                                <div className="player-info">
                                    <h1>
                                        <span className="summoner-name">{accountInfo.gameName} #{accountInfo.tagLine}</span>
                                    </h1>
                                    <div className="update">
                                        <button className="update-button">Update</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div className="some-divider">
                            <div style={{marginLeft: "20px"}}>Recent Games</div>
                        </div>
                    </div>
                    <div className="content-container">
                        {renderMatches()}
                    </div>
                </>
            )}
        </div>
    )
}