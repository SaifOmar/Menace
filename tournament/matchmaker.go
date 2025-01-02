package tournament

import (
	// "fmt"
	"math"
	"math/rand"

	"TournamentProject/helpers"

	"TournamentProject/match"
	"TournamentProject/player"
)

type MatchMaker struct {
	playersPool   []*player.Player
	mS            []*player.Player
	k             int
	round         int
	RoundFinished bool
	Logger        *helpers.TournamentLogger
}

func NewMatchMaker(players []*player.Player, logger *helpers.TournamentLogger) *MatchMaker {
	mM := &MatchMaker{
		playersPool:   players,
		mS:            make([]*player.Player, len(players)),
		k:             100,
		round:         0,
		RoundFinished: false,
		Logger:        logger,
	}
	return mM
}

func QuickSort(players *[]*player.Player, low, high int) {
	if low < high {
		pi := partition(players, low, high)
		QuickSort(players, low, pi-1)
		QuickSort(players, pi+1, high)
	}
}

func partition(players *[]*player.Player, low, high int) int {
	pivot := (*players)[high].AdjustedElo
	i := low - 1
	for j := low; j <= high-1; j++ {
		// If current player has Elo less than or equal to pivot, swap them
		if (*players)[j].AdjustedElo <= pivot {
			i++
			(*players)[i], (*players)[j] = (*players)[j], (*players)[i]
		}
	}
	// Swap the pivot with the element at i+1
	(*players)[i+1], (*players)[high] = (*players)[high], (*players)[i+1]
	return i + 1
}

// returns array of 2 players that have the lowest Elo diff to play the next match
// ended up not using it
func minDiffThatShit(pSlice *[]*player.Player) []*player.Player {
	s := *pSlice
	l := len(s)

	QuickSort(pSlice, 0, l-1)

	minDiff := math.MaxInt
	var res [][]*player.Player

	for i := 0; i < l-1; i++ {
		diff := s[i+1].AdjustedElo - s[i].AdjustedElo
		if diff < minDiff {
			minDiff = diff
			tempArr := []*player.Player{s[i], s[i+1]}
			res = [][]*player.Player{tempArr}
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res[0]
}

func adjustMeStepBro(slice *[]*player.Player) {
	s := *slice
	for _, p := range s {
		adjustedWp := ((p.WP/5.0 - 10.0) / 100.0 * 2.0 * float64(p.Elo)) + float64(p.Elo)
		p.AdjustedElo = int(adjustedWp)

	}
}

func (matchMaker *MatchMaker) MakeMatch() (*match.Match, *MatchMaker) {
	if matchMaker.round == 0 {
		matchMaker.Logger.Info("First Round Has started")
		m := firstRound(matchMaker)
		return m, matchMaker
	}
	if matchMaker.RoundFinished {
		MS := make([]*player.Player, len(matchMaker.playersPool))
		copy(MS, matchMaker.playersPool)
		matchMaker.mS = MS
		matchMaker.RoundFinished = false
	}

	adjustMeStepBro(&matchMaker.mS)
	QuickSort(&matchMaker.mS, 0, len(matchMaker.mS)-1)
	ps := matchMaker.mS

	m := match.NewMatch([2]*player.Player{ps[0], ps[1]}, matchMaker.Logger)
	count := 0
	for i := len(matchMaker.mS) - 1; i >= 0; i-- {
		if count == 2 {
			break
		}
		if matchMaker.mS[i].Name == ps[0].Name || matchMaker.mS[i].Name == ps[1].Name {
			count++
			helpers.DeleteSliceElement(&matchMaker.mS, i)
		}
	}
	calculateElo(m, matchMaker.k)
	calculateWinPercentage(m)
	if len(matchMaker.mS) == 0 {
		matchMaker.RoundFinished = true
	}

	return m, matchMaker
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

func firstRound(mM *MatchMaker) *match.Match {
	mM.Logger.Info("First Round Match")
	length := copy(mM.mS, mM.playersPool)

	n1 := Random(length)
	n2 := Random(length)
	for n2 == n1 {
		n2 = Random(length)
	}
	p1 := mM.mS[n1]
	p2 := mM.mS[n2]

	players := [2]*player.Player{p1, p2}
	m := match.NewMatch(players, mM.Logger)

	calculateElo(m, mM.k)
	calculateWinPercentage(m)
	if n1 > n2 {
		helpers.DeleteSliceElement(&mM.mS, n1)
		helpers.DeleteSliceElement(&mM.mS, n2)
	} else {
		helpers.DeleteSliceElement(&mM.mS, n2)
		helpers.DeleteSliceElement(&mM.mS, n1)
	}

	if len(mM.mS) == 0 {
		mM.round = 1
		mM.RoundFinished = true
		return m
	} else {
		firstRound(mM)
	}

	return m
}

func Random(n int) int {
	if n == 1 {
		n++
	}
	return rand.Intn(n)
}
