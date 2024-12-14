package main

import (
	// "fmt"

	"TournamentProject/match"
	"TournamentProject/player"
)

func main() {
	Salah := new(player.Player)
	Yousry := new(player.Player)
	Salah.NewPlayer("Salah", 193, 120, 5, 2, -120)
	Yousry.NewPlayer("Yousry", 178, 74, 2, 4, -190)
	// fmt.Println(Salah.GetAbility())
	// fmt.Println(Salah.GetPassive())
	Match := new(match.Match)
	Match.NewMatch(*Salah, *Yousry)
	Match.FirstBlow()
	Match.FinishMatch()
}
