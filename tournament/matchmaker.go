package tournament

import (
	"fmt"
	"math"

	"TournamentProject/helpers"

	"TournamentProject/match"
	"TournamentProject/player"
)

type MatchMaker struct {
	playersPool []*player.Player
	// k is the factor by how much Elo will be changed
	// TODO: should be dynamic
	roundMatch int
	round      int
	Logger     *helpers.TournamentLogger
}

func NewMatchMaker(players []*player.Player, logger *helpers.TournamentLogger) *MatchMaker {
	matchMaker := &MatchMaker{
		playersPool: players,
		round:       0,
		Logger:      logger,
	}
	return matchMaker
}

func adjustMeStepBro(slice *[]*player.Player) {
	s := *slice
	for _, p := range s {
		// is how elo is adjusted based on wp
		// adjustedwp is used to match players of similar skill level
		adjustedWp := ((p.WP - 50.0) / 10.0) * math.Pow(float64(p.Elo), 0.65)
		p.AdjustedElo = int(adjustedWp)

	}
}

func (matchMaker *MatchMaker) MakeMatch() *match.Match {
	m := matchMaker.roundStart()
	matchMaker.roundEnd(m)

	if matchMaker.roundMatch == len(matchMaker.playersPool)/2 {
		matchMaker.round++
	}

	return m
}

func calculateWinPercentage(m *match.Match) {
	for _, p := range m.Players {
		p.WP = float64(p.WinCount) * 100 / float64(p.NMatches)
	}
}

func calculateElo(m *match.Match, k int) {
	m.Winner.Elo += k
	m.Winner.WinCount++
	m.Players[1].NMatches += 1
	m.Players[0].NMatches += 1
	if m.Winner == m.Players[0] {
		m.Players[1].Elo -= k / 2
	} else {
		m.Players[0].Elo -= k / 2
	}
}

func firstRound(matchMaker *MatchMaker) *match.Match {
	// this might change later to have all logs be controller from tournament
	matchMaker.Logger.Info("First Round Match")

	length := copy(matchMaker.playersPool, matchMaker.playersPool)

	matchMaker.roundMatch++
	n1 := helpers.Random(length)
	n2 := helpers.Random(length)
	for n2 == n1 {
		n2 = helpers.Random(length)
	}
	p1 := matchMaker.playersPool[n1]
	p2 := matchMaker.playersPool[n2]

	players := [2]*player.Player{p1, p2}
	m := match.NewMatch(players, matchMaker.Logger)

	matchMaker.roundEnd(m)

	if matchMaker.roundMatch == len(matchMaker.playersPool)/2 {
		matchMaker.round = 1
		return m
	} else {
		firstRound(matchMaker)
	}

	return m
}

func (matchMaker *MatchMaker) roundStart() *match.Match {
	if matchMaker.round == 0 {
		matchMaker.Logger.Info("First Round Has started")
		m := firstRound(matchMaker)
		return m
	}
	matchMaker.roundMatch++
	adjustMeStepBro(&matchMaker.playersPool)

	for _, p := range matchMaker.playersPool {
		matchMaker.Logger.Debug(fmt.Sprintf("Name : %s, Wp : %f , Elo : %d, Adjsted : %d", p.Name, p.WP, p.Elo, p.AdjustedElo))
	}

	helpers.QuickSort(&matchMaker.playersPool, 0, len(matchMaker.playersPool)-1)

	player1, player2 := helpers.FindCandidates(matchMaker.playersPool)
	m := match.NewMatch([2]*player.Player{player1, player2}, matchMaker.Logger)
	return m
}

func (matchMaker *MatchMaker) roundEnd(m *match.Match) {
	calculateElo(m, int(matchMaker.calculateKfactor(m)))
	calculateWinPercentage(m)
}

func (matchMaker *MatchMaker) calculateKfactor(m *match.Match) float32 {
	// this is how the k factor is calculated
	// this will be changed later
	k := 0.0
	for _, p := range m.Players {
		if p.Elo < 2400 {
			k = 32.0
		} else if p.Elo >= 2400 && p.Elo < 3000 {
			k = 24.0
		} else {
			k = 16.0
		}
	}
	m.SetKfactor(float32(k))
	return float32(k)
}

// WARNING: claude made this I don't know what is this

// func (matchMaker *MatchMaker) calculateKfactors(m *match.Match) {
// 	// Calculate base K-factors for each player individually
// 	for i, player := range m.Players {
// 		// Base K-factor based on player's rating
// 		var baseK float32
// 		switch {
// 		case player.Elo < 2400:
// 			baseK = 32.0
// 		case player.Elo < 3000:
// 			baseK = 24.0
// 		default:
// 			baseK = 16.0
// 		}
//
// 		// Adjust K-factor based on number of games played (uncertainty factor)
// 		// Players with fewer games have higher K-factors to help them reach their true skill level faster
// 		gamesPlayedFactor := float32(1.0)
// 		if player.NMatches < 30 {
// 			gamesPlayedFactor = 1.0 + float32(30-player.NMatches)/30.0
// 		}
//
// 		// Adjust K-factor based on win/loss streaks
// 		// Players on streaks might be improving/declining rapidly
// 		streakFactor := float32(1.0)
// 		if player.WinStreak > 3 {
// 			streakFactor = 1.0 + (float32(player.WinStreak) * 0.05)
// 		} else if player.LoseStreak > 3 {
// 			streakFactor = 1.0 + (float32(player.LoseStreak) * 0.05)
// 		}
//
// 		// Adjust K-factor based on performance consistency
// 		// More volatile players get slightly higher K-factors
// 		consistencyFactor := float32(1.0)
// 		if player.RatingDeviation > 100 {
// 			consistencyFactor = 1.0 + (player.RatingDeviation-100)/200
// 		}
//
// 		// Calculate final K-factor with all adjustments
// 		finalK := baseK * gamesPlayedFactor * streakFactor * consistencyFactor
//
// 		// Enforce reasonable boundaries for K-factor
// 		if finalK < 8.0 {
// 			finalK = 8.0 // Minimum K-factor
// 		} else if finalK > 64.0 {
// 			finalK = 64.0 // Maximum K-factor
// 		}
//
// 		// Store individual K-factor for this player
// 		m.Players[i].KFactor = finalK
// 	}
//
// 	// Calculate match's overall K-factor (average of all players) for compatibility
// 	totalK := float32(0)
// 	for _, player := range m.Players {
// 		totalK += player.KFactor
// 	}
//
// 	avgK := float32(32.0) // Default if no players
// 	if len(m.Players) > 0 {
// 		avgK = totalK / float32(len(m.Players))
// 	}
//
// 	// Set match's overall K-factor
// 	m.SetKfactor(avgK)
// }
