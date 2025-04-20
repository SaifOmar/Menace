package match

import (
	"fmt"
	"math/rand"

	"TournamentProject/helpers"
	"TournamentProject/luck"
	"TournamentProject/player"
)

type Match struct {
	Players [2]*player.Player
	Winner  *player.Player
	k       float32
	*luck.Luck
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
		Luck:     luck.NewLuck(),
	}
	logger.Info(fmt.Sprintf("%s is facing %s", players[0].Name, players[1].Name))
	M.FirstBlow()
	return M
}

func (match *Match) SetKfactor(k float32) {
	match.k = k
}

func (match *Match) GetKfactor() float32 {
	return match.k
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
	match.endMatch()
	return match
}

func (match *Match) endMatch() {
	match.Logger.Info("Match is over")
	match.Logger.Info(fmt.Sprintf("%s has won the match", match.Winner.Name))
	for _, p := range match.Players {
		p.Hp = 100
	}
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
	luck := match.GetLucky()
	if luck.Value != "Unlucky" {
		match.Logger.Debug("Someone just got Lucky")
		match.applyEffect(luck)
	}
	hitter.Hit(hittee)
	hitter.AbilityHit(hittee)
	match.checkForEnd()
	match.Logger.Info(fmt.Sprintf("%s has hit %s and his hp is %d", hitter.Name, hittee.Name, hittee.Hp))
}

func (match *Match) applyEffect(luck *luck.Luck) {
	target := match.Players[luck.Player]
	target.Hp -= luck.Damage //-20
	match.Logger.Debug(luck.Note)
}

func (match *Match) checkForEnd() {
	for _, p := range match.Players {
		if p.Hp <= 0 {
			match.Finished = true
			match.Winner = match.Players[1]
			if match.Winner == p {
				match.Winner = match.Players[0]
			}
			break
		}
	}
}
