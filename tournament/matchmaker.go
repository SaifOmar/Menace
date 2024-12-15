package tournament

import (
	"math"

	"TournamentProject/match"
	"TournamentProject/player"
)

type MatchMaker struct {
	playersPool []*player.Player
}

func NewMatchMaker(players []*player.Player) (*MatchMaker, *match.Match) {
	mM := &MatchMaker{
		playersPool: players,
	}
	match := mM.MakeMatch()
	return mM, match
}

func (matchMaker *MatchMaker) MakeMatch() *match.Match {
	pool := matchMaker.playersPool
	for i, p := range pool {

		currWp, _ := p.CalculateWinPercentage()
		nextWp, _ := pool[i+1].CalculateWinPercentage()

		wpDiff := math.Abs(currWp - nextWp)

		if wpDiff <= 10.0 {
			m := match.NewMatch([2]*player.Player{p, pool[i+1]})
			return m
		}

	}

	return nil
}
