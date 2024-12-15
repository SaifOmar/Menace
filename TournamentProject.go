package main

import (
	"fmt"
	// "fmt"

	"TournamentProject/player"
	"TournamentProject/tournament"
)

func main() {
	players := player.CreatePlayers()
	tournament, _ := tournament.NewTournament(players)

	if tournament.Winner.Name != "" {
		fmt.Println(tournament.Winner.Name)
	}
}
