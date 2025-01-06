package match

import (
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
	match = match.runMatch()
	match.Logger.Info("Match is over")
	match.Logger.Info(fmt.Sprintf("%s has won the match", match.Winner.Name))
	for _, p := range match.Players {
		p.Hp = 100
	}
	return match
}

func (match *Match) randomHit() *Match {
	n := helpers.Random(2)
	if n == 1 {
		match.matchHit(match.Players[n], match.Players[0])
	} else {
		match.matchHit(match.Players[n], match.Players[1])
	}
	return match
}

func (match *Match) runMatch() *Match {
	for !match.Finished {
		match = match.randomHit()
	}

	return match
}

func (match *Match) matchHit(hitter *player.Player, hittee *player.Player) {
	hitter.Hit(hittee)
	hitter.AbilityHit(hittee)
	s := checkPlayerHp(hittee)
	if s == "dead" {
		match.Finished = true
		match.Winner = hitter
	}
	match.Logger.Info(fmt.Sprintf("%s has hit %s and his hp is %d", hitter.Name, hittee.Name, hittee.Hp))
}

func checkPlayerHp(p *player.Player) string {
	if p.Hp == 0 {
		return "dead"
	}
	return "alive"
}
