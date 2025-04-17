package main

import (
	"fmt"

	"TournamentProject/helpers"
	"TournamentProject/tournament"
)

func main() {
	logger := helpers.NewTournamentLogger(true, "tournament.log")
	logger.EnableDebug(helpers.Debug)
	t := tournament.NewTournament(logger)
	// count which I don't remember why I put it here
	co := 0
	for _, p := range t.Players {
		fmt.Println("i, name, wp , elo, adjusted elo: ", co, p.Name, p.WP, p.Elo, p.AdjustedElo)
		co++
	}
	fmt.Println(t.Winner.Name, "has won the tourny")
	// t.Logger.GetLogs()
	// if t.Winner != nil {
	// 	fmt.Println(t.Winner.Name)
	// }
}
