package match

import (
	// "fmt"
	"fmt"
	"math/rand"

	"TournamentProject/helpers"
	"TournamentProject/player"
)

type Match struct {
	Players  [2]*player.Player
	Winner   *player.Player
	Duration int
	Finished bool
	Logger   *helpers.TournamentLogger
}

func (match *Match) matchHit(hitter *player.Player, hittee *player.Player) {
	match.Logger.Info(fmt.Sprintf("%s has hit %s", hitter.Name, hittee.Name))
	hitter.Hit(hittee)
	hitter.AbilityHit(hittee)
}

func NewMatch(players [2]*player.Player, logger *helpers.TournamentLogger) *Match {
	M := &Match{
		Players:  players,
		Finished: false,
		Duration: 10,
		Logger:   logger,
	}
	logger.Info(fmt.Sprintf("%s is facing %s", players[0].Name, players[1].Name))
	M.FirstBlow()
	return M
}

func (match *Match) FirstBlow() *Match {
	match.Logger.Info("First Blow 🤜")
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
	for !match.Finished {
		if match.Players[0].Hp == 0 {
			match.Finished = true
			match.Winner = match.Players[0]
			return match
		} else if match.Players[1].Hp == 0 {
			match.Finished = true
			match.Winner = match.Players[1]
		}
		match.matchHit(match.Players[0], match.Players[1])
		match.matchHit(match.Players[1], match.Players[0])
	}
	match.Logger.Info("Match is over")
	match.Logger.Info(fmt.Sprintf("%s has won the match", match.Winner.Name))
	return match
}
