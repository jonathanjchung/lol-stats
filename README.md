# lol stats (op.gg clone)
op.gg clone as a means to learn Go


## getting started
will reformat this later

under root directory
* npm run install-all
* npm run setup
* npm run dev

### env variables (under /server)
* ```riot_api_key``` - go to developer.riotgames.com
* ```user, password, database, host``` - fill these out (if user isn't postgres, go to /server/package.json and adjust "setup" script)

### todo
* stop wrapping everything in divs
* update button functionality on summoner page
* region dropdown on search page, implement more than NA support, only NA1 tagline works
* fix case sensitivity in database and summoner search
* utilize datadragon to replace ids with icons
* ranked leaderboard
* summoner page display more information (rank, champ winrates, etc)
