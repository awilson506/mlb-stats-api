# Golang MLB Stats API 
A small api wrapper service for getting stats about teams using dates and teamIds.

## Getting the service setup & running
If you don't already have go installed you can download it and install it from here [Download Golang](https://go.dev/doc/install)

Note this application requires the latest go version `1.19` but supports back to `1.17` if you already have a previous version
installed.  You will have to update the go module to require a later version if you choose not to update.  This can be done by 
running: 
```sh
go mod edit -go=1.MY_VERSION
```

## Start the server
In the root directory run:
```sh
go run cmd/server/main.go
```
The server will default to port 8080, you can specify a different port as well with the `port` param:
In the root directory run:
```sh
go run cmd/server/main.go --port=:8081
```

## Using the api
This application offers 1 endpoint:

### Get MLB Game Stats
Takes in the id of the MLB team and the date in which the games were played on.  The API should take
the team id and order order them at the top of the JSON object according to the game status and order played in
```
curl http://localhost:8080/v1/stats?team-id=147&date=2022-07-21
```
Trimmed Example Response:
```json
{
    "copyright": "Copyright 2023 MLB Advanced Media, L.P.  Use of any content on this page acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt",
    "totalItems": 6,
    "totalEvents": 0,
    "totalGames": 6,
    "totalGamesInProgress": 0,
    "dates": [
        {
            "date": "2022-07-21",
            "totalItems": 6,
            "totalEvents": 0,
            "totalGames": 6,
            "totalGamesInProgress": 0,
            "games": [
                {
                    "gamePk": 662776,
                    "link": "/api/v1.1/game/662776/feed/live",
                    "gameType": "R",
                    "season": "2022",
                    "gameDate": "2022-07-21T17:10:00Z",
                    "officialDate": "2022-07-21",
                    "status": {
                        "abstractGameState": "Final",
                        "codedGameState": "F",
                        "detailedState": "Final",
                        "statusCode": "F",
                        "startTimeTBD": false,
                        "abstractGameCode": "F"
                    },
                    "teams": {
                        "away": {
                            "leagueRecord": {
                                "wins": 64,
                                "losses": 29,
                                "pct": ".688"
                            },
                            "score": 2,
                            "team": {
                                "id": 147,
                                "name": "New York Yankees",
                                "link": "/api/v1/teams/147"
                            },
                            "isWinner": false,
                            "splitSquad": false,
                            "seriesNumber": 31
                        },
                        "home": {
                            "leagueRecord": {
                                "wins": 60,
                                "losses": 32,
                                "pct": ".652"
                            },
                            "score": 3,
                            "team": {
                                "id": 117,
                                "name": "Houston Astros",
                                "link": "/api/v1/teams/117"
                            },
                            "isWinner": true,
                            "splitSquad": false,
                            "seriesNumber": 31
                        }
                    },
                    "venue": {
                        "id": 2392,
                        "name": "Minute Maid Park",
                        "link": "/api/v1/venues/2392"
                    },
                    "content": {
                        "link": "/api/v1/game/662776/content"
                    },
                    "isTie": false,
                    "gameNumber": 1,
                    "publicFacing": true,
                    "doubleHeader": "S",
                    "gamedayType": "P",
                    "tiebreaker": "N",
                    "calendarEventID": "14-662776-2022-07-21",
                    "seasonDisplay": "2022",
                    "dayNight": "day",
                    "description": "Rescheduled from 4/5",
                    "scheduledInnings": 9,
                    "reverseHomeAwayStatus": false,
                    "inningBreakLength": 120,
                    "gamesInSeries": 2,
                    "seriesGameNumber": 1,
                    "seriesDescription": "Regular Season",
                    "recordSource": "S",
                    "ifNecessary": "N",
                    "ifNecessaryDescription": "Normal Game"
                         }
            ],
            "events": []
        }
    ]
}         
```

## Assumptions made
- The `statusCode` of a `Live` game is `I`
- Double headers should always have a `doubleheader` status of `Y` or `S`
    - Type `Y` games will have `startTImeTBD` set to true for the second game
    - Type `S` games will have both `gameDate` fields populated 
- If no games with the specified team Id are found the remaining list will be returned

## TODO:
See here for [todo](todo.md) doc