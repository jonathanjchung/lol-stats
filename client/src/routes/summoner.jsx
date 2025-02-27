import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import Navbar from "../components/navbar";
import MatchData from "../components/matchdata";
import profileIcon from "../icons/opgg.jpg";
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
                <div>
                    {matchHistory.map((m) => (
                        <MatchData key={m.info.gameId} match={m} summName={accountInfo.gameName} />
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
                        <div className="header-wrapper">
                            <div className="profile-icon">
                                <img src={profileIcon} alt="Profile icon"/>
                                <div>{summonerInfo.summonerLevel}</div>
                            </div>
                            <div className="player-info">
                                <div>{accountInfo.gameName} #{accountInfo.tagLine}</div>
                            </div>
                        </div>
                    </div>
                    <div className="content-container">
                        <h2>Recent games</h2>
                        {renderMatches()}
                    </div>
                </>
            )}
        </div>
    )
}