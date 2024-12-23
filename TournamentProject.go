package main

import (
	// "fmt"
	"fmt"

	// "TournamentProject/player"
	"TournamentProject/tournament"
)

func main() {
	t := tournament.NewTournament()
	co := 0
	for _, p := range t.Players {

		fmt.Println("i, name, wp , elo: ", co, p.Name, p.WP, p.Elo)
		co++
	}
	// if t.Winner != nil {
	// 	// fmt.Println(t.Winner.Name)
	// }
}
