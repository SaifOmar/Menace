package match

import (
	// "fmt"
	// "math/rand"

	"TournamentProject/player"
)

type Match struct {
	Player1  *player.Player
	Player2  *player.Player
	Winner   player.Player
	Duration int
	Finished bool
}

func (match *Match) NewMatch(player1 player.Player, player2 player.Player) *Match {
	match.Player1 = &player1
	match.Player2 = &player2
	match.Finished = false
	match.Duration = 10
	return match
}

func (match *Match) FirstBlow() *Match {
	// rN := rand.Intn(2)
	// if rN == 1 {
	// 	match.Player2.Hit(match.Player1)
	// 	fmt.Print("rn: ")
	// 	fmt.Println(rN)
	// 	return match
	// }
	// match.Player1.Hit(match.Player2)
	// fmt.Print("rn: ")
	// fmt.Println(rN)
	return match
}

func (match *Match) matchHit(hitter *player.Player, hittee *player.Player) {
	hitter.Hit(hittee)
}

func (match *Match) FinishMatch() *Match {
	for i := 0; i < 33; i++ {
		if match.Player1.Hp == 0 || match.Player2.Hp == 0 {
			match.Finished = true
			return match
		}
		match.matchHit(match.Player1, match.Player2)
		match.matchHit(match.Player2, match.Player1)
	}
	return match
}
