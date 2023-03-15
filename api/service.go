package api

import (
	"log"
	"os"
)

// gameLiveStatus - statusCode will be "I" when the game is live it's likely these
// status' could be kept in a db if this were in production: I, P, S...
const gameLiveStatus = "I"

type Client interface {
	GetFavoriteMLBStats(favoriteTeamId int, date string) GameStats
}

type client struct {
	logger *log.Logger
}

// New - get a new instance of the client
func New() Client {
	return &client{
		logger: log.New(os.Stdout, "", 5),
	}
}

// GetFavoriteMLBStats - get a list of mlb games with the "favorite" teams listed first by teamId
func (c *client) GetFavoriteMLBStats(favoriteTeamId int, date string) GameStats {
	gameStats := c.GetMLBEventsByDate(date)
	favoriteGames := make(map[int]Game)
	gameCount := 0

	// theres only one or no games
	if len(gameStats.Dates) == 0 || len(gameStats.Dates[0].Games) == 1 {
		return gameStats
	}

	for i := 0; i < len(gameStats.Dates[0].Games); i++ {
		if gameStats.Dates[0].Games[i].Teams.Away.Team.ID == favoriteTeamId || gameStats.Dates[0].Games[i].Teams.Home.Team.ID == favoriteTeamId {
			// remove the games and build a list of games/DH games to sort later if need be
			favoriteGames[gameCount] = gameStats.Dates[0].Games[i]
			gameStats.Dates[0].Games = c.removeGames(gameStats.Dates[0].Games, i)
			i--
			gameCount++
		}
	}

	if gameCount > 1 {
		// sort the double header games
		favoriteGames = c.sortGames(favoriteGames)
	}

	// add the games/game back to the the begining of the list
	gameStats.Dates[0].Games = c.prependGames(gameStats.Dates[0].Games, favoriteGames)

	return gameStats
}

// sortGames - order games based on there status and or gameDates
func (c *client) sortGames(games map[int]Game) map[int]Game {

	// check if the games are live
	if games[0].Status.StatusCode == gameLiveStatus {
		return games
	}

	if games[1].Status.StatusCode == gameLiveStatus {
		c.swap(games)
		return games
	}

	// check which type of DH games they are and sort
	switch gameType := games[0].DoubleHeader; gameType {
	case "Y":
		if games[0].Status.StartTimeTBD {
			c.swap(games)
		}
	case "S":
		if games[0].GameDate.After(games[1].GameDate) {
			c.swap(games)
		}
	}
	return games
}

// swap - swap the game order
func (c *client) swap(games map[int]Game) {
	games[0], games[1] = games[1], games[0]
}

// removeGames - remove a game from the games list
func (c *client) removeGames(games []Game, index int) []Game {
	return append(games[:index], games[index+1:]...)
}

// prependGames - add the sorted games to the begining of the array
func (c *client) prependGames(games []Game, favoriteGames map[int]Game) []Game {
	// start from the last spot of to preserve order
	for i := len(favoriteGames) - 1; i >= 0; i-- {
		games = append([]Game{favoriteGames[i]}, games...)
	}
	return games
}
