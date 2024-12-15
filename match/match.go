package match

import (
	// "fmt"
	"math/rand"

	"TournamentProject/player"
)

type Match struct {
	Players  [2]*player.Player
	Winner   player.Player
	Duration int
	Finished bool
}

func (match *Match) matchHit(hitter *player.Player, hittee *player.Player) {
	hitter.Hit(hittee)
	hitter.AbilityHit(hittee)
}

func NewMatch(players [2]*player.Player) *Match {
	M := &Match{
		Players:  players,
		Finished: false,
		Duration: 10,
	}
	M.FirstBlow()
	return M
}

func (match *Match) FirstBlow() *Match {
	rN := rand.Intn(2)
	if rN == 1 {
		match.matchHit(match.Players[1], match.Players[0])
		match.FinishMatch()
		return match
	}
	match.matchHit(match.Players[0], match.Players[1])
	match.FinishMatch()

	return match
}

func (match *Match) FinishMatch() *Match {
	for i := 0; i < 33; i++ {
		if match.Players[0].Hp == 0 {
			match.Finished = true
			match.Winner = *match.Players[0]
			return match
		} else if match.Players[1].Hp == 0 {
			match.Finished = true
			match.Winner = *match.Players[1]
		}
		match.matchHit(match.Players[0], match.Players[1])
		match.matchHit(match.Players[1], match.Players[0])
	}
	return match
}
