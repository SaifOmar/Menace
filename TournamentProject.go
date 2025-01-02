package main

import (
	// "fmt"
	"fmt"

	// "TournamentProject/player"
	"TournamentProject/helpers"
	"TournamentProject/tournament"
)

func main() {
	logger := helpers.NewTournamentLogger(true, "tournament.log")
	t := tournament.NewTournament(logger)
	co := 0
	for _, p := range t.Players {

		fmt.Println("i, name, wp , elo: ", co, p.Name, p.WP, p.Elo)
		co++
	}
	t.Logger.GetLogs()
	// if t.Winner != nil {
	// 	// fmt.Println(t.Winner.Name)
	// }
}
