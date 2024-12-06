import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import Navbar from "../components/navbar";

export default function Summoner() {
    const { summonerName, tagLine } = useParams();
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
                    {matchHistory.map((match) => {
                        const summonerData = findSummonerData(match, summonerName);
                        return (
                            <div id="match" key={match.metadata.matchId} style={{border: "1px solid", margin: "10px", padding: "10px", backgroundColor: "white"}}>
                                <div style={{display: "flex", gap: "10px"}}>
                                    <p><strong>Match Id: </strong>{match.metadata.matchId}</p>
                                    <p><strong>Gamemode: </strong>{match.info.gameMode}</p>
                                    <p><strong>Duration: </strong>{match.info.gameDuration}</p>
                                    <p><strong>Outcome: </strong>{summonerData ? (summonerData.win ? "Victory" : "Defeat") : "N/A"}</p>
                                </div>
                                <div>
                                    <p><strong>Champion: </strong>{summonerData ? summonerData.championName : "Not Found"}</p>
                                    <p><strong>Lane: </strong>{summonerData ? summonerData.individualPosition : "N/A"}</p>
                                </div>
                            </div>
                        )
                    })}
                </div>
            )
        }
    }

    const findSummonerData = (match, summoner) => {
        return match.info.participants.find(
            (participant) => participant.summonerName === summoner
        );
    }

    useEffect(() => {
        console.log(`Looking at summoner: ${summonerName}#${tagLine}`);
        setLoading(true);
        fetch(url)
        .then((res) => {
            if (!res.ok) {
                throw new Error('http error: ' + res.status);
            }
            return res.json();
        })
        .then((data) => {
            console.log(data.summoner);
            console.log(data.matches)
            setSummonerInfo(data.summoner);
            setMatchHistory(data.matches);
            setLoading(false);
        })
        .catch((err) => {
            console.error(err);
            setError(err);
            setLoading(false);
        });
    }, [url]);

    return (
        <div>
            <Navbar />
            {loading && <div>Loading...</div>}
            {error && <div>Error: {error.message}</div>}
            
            {!loading && !error && (
                <>
                    <h1>{summonerName} #{tagLine}</h1>
                    <div>Summoner Level: {summonerInfo.summonerLevel}</div>
                    <div>Profile Icon: {summonerInfo.profileIconId}</div>
                    <h2>Recent games</h2>
                    {renderMatches()}
                </>
            )}
        </div>
    )
}