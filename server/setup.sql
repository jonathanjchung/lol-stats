DROP DATABASE IF EXISTS opgg;
CREATE DATABASE opgg;
\c opgg

CREATE TABLE player_info (
    puuid text PRIMARY KEY,
    summoner_name text NOT NULL,
    tag_line text NOT NULL,
    account jsonb NOT NULL,
    summoner jsonb NOT NULL,
    matches jsonb NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (summoner_name, tag_line)
);