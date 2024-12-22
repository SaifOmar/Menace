package main

import (
	"fmt"
	// "fmt"

	// "TournamentProject/player"
	"TournamentProject/tournament"
)

func main() {
	t := tournament.NewTournament()
	if t.Winner != nil {
		fmt.Println(t.Winner.Name)
	}
}
