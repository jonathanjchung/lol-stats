import './matchdata.css'

function findSummonerObjectByName(array, name) {
    return array.find((summoner) => {
        return summoner.riotIdGameName === name;
    })
}

function getKDA(kills, deaths, assists) {
    var kda = (kills + assists) / deaths;
    var roundedKda = kda.toFixed(2);
    var output = [];

    if (deaths == 0) {
        output.push("Perfect");
    } else {
        output.push(roundedKda, ":", "1 KDA");
    }

    return output;
}

function secondsToMMSS(time) {
    const minutes = Math.floor(time / 60);
    const seconds = time - (minutes * 60);
    var output = [];

    output.push(minutes, "m ", seconds, "s");

    return output;
}

const MatchData = ({ match, summName }) => {
    const summonerData = findSummonerObjectByName(match.info.participants, summName);

    return (
        <div className="contents" key={match.metadata.matchId} style={{backgroundColor: summonerData.win ? "#ecf2ff" : "#fff1f3"}}>
            <div className="match-overview">
                <p><strong>{match.info.gameMode}</strong></p>
                <p><strong style={{color: summonerData.win ? "#4171d6" : "#d31a45"}}>{summonerData ? (summonerData.win ? "Victory" : "Defeat") : "N/A"}</strong></p>
                <p><strong>{secondsToMMSS(match.info.gameDuration)}</strong></p>
            </div>
            <div className="game-details">
                <div className="top-row">
                    <div className="champ-and-loadout">
                        <div className="champ">
                            <p><strong>{summonerData ? summonerData.championName : "Not Found"}</strong></p>
                        </div>
                        <div className="loadout">
                            <ul className="sums">
                                <li className="item-and-sum image-icon">{summonerData.summoner1Id}</li>
                                <li className="item-and-sum image-icon">{summonerData.summoner2Id}</li>
                            </ul>
                            <ul className="runes">
                                <li className="rune-icon image-icon">{summonerData.perks.styles[0].selections[0].perk}</li>
                                <li className="rune-icon image-icon">{summonerData.perks.styles[0].style}</li>
                            </ul>
                        </div>
                    </div>
                    <div className="kda">
                        <div className="scoreline">
                            <strong>{summonerData.kills}</strong>
                            /
                            <strong>{summonerData.deaths}</strong>
                            /
                            <strong>{summonerData.assists}</strong>
                        </div>
                        <div className="kda-calculated">
                            {getKDA(summonerData.kills, summonerData.deaths, summonerData.assists)}
                        </div>
                    </div>
                </div>
                <div className="bottom-row">
                    <ul className="item-list">
                        <li className="item-and-sum image-icon">{summonerData.item0}</li>
                        <li className="item-and-sum image-icon">{summonerData.item1}</li>
                        <li className="item-and-sum image-icon">{summonerData.item2}</li>
                        <li className="item-and-sum image-icon">{summonerData.item3}</li>
                        <li className="item-and-sum image-icon">{summonerData.item4}</li>
                        <li className="item-and-sum image-icon">{summonerData.item5}</li>
                        <li className="item-and-sum image-icon">{summonerData.item6}</li>
                    </ul>
                </div>
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