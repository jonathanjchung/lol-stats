function findSummonerObjectByName(array, name) {
    return array.find((summoner) => {
        return summoner.riotIdGameName === name;
    })
}

const MatchData = ({ match, summName }) => {
    const summonerData = findSummonerObjectByName(match.info.participants, summName)

    return (
        <div className="contents" key={match.metadata.matchId} style={{border: "1px solid", marginBottom: "10px", padding: "10px", backgroundColor: "white", display: "flex", flexDirection: "row"}}>
            <div className="match-details">
                <p><strong>Gamemode: </strong>{match.info.gameMode}</p>
                <p><strong>Duration: </strong>{match.info.gameDuration}</p>
                <p><strong>Outcome: </strong>{summonerData ? (summonerData.win ? "Victory" : "Defeat") : "N/A"}</p>
            </div>
            <div className="champ-details">
                <p><strong>Champion: </strong>{summonerData ? summonerData.championName : "Not Found"}</p>
                <p><strong>Lane: </strong>{summonerData ? summonerData.individualPosition : "N/A"}</p>
            </div>
            <div className="participants">
                <div className="blue-side">
                    {match.info.participants.slice(0, 5).map((p) => <div key={`${p.puuid || p.riotIdGameName}`}>{p.riotIdGameName || "unknown"}</div>)}
                </div>
                <div className="red-side">
                    {match.info.participants.slice(5).map((p) => <div key={`${p.puuid || p.riotIdGameName}`}>{p.riotIdGameName || "unknown"}</div>)}
                </div>
            </div>
        </div>
    );
};

export default MatchData;