package tournament

import (
	// "fmt"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"

	"TournamentProject/helpers"

	"TournamentProject/match"
	"TournamentProject/player"
)

type MatchMaker struct {
	playersPool []*player.Player
	seen        map[string]int // keeps track of hwo many times playes have played
	k           int
	round       int
	playerQueue *helpers.PlayerQueue
}

func NewMatchMaker(players []*player.Player) *MatchMaker {
	mM := &MatchMaker{
		playersPool: players,
		seen:        make(map[string]int, len(players)),
		k:           100,
		round:       0,
		playerQueue: &helpers.PlayerQueue{},
	}
	return mM
}

func (matchMaker *MatchMaker) MakeMatch() (*match.Match, *MatchMaker) {
	if matchMaker.round == 0 {
		m := firstRound(matchMaker)
		return m, matchMaker
	}
	// var elos []int
	var adjustedElos []int
	eloDB := make(map[int]int)
	for i, p := range matchMaker.playersPool {
		curr := p.Elo
		wp := int(p.WinPercentage)
		adjust := curr + (((wp/5 - 10) / 100) * 2 * curr)
		adjustedElos = append(adjustedElos, adjust)
		// elos = append(elos, p.Elo)
		eloDB[adjust] = i
	}
	fmt.Println(adjustedElos, matchMaker.playersPool)

	bestEloMatch := findBestEloMatch(adjustedElos)

	fmt.Println("indexes")
	fmt.Println(adjustedElos)
	p1Index := eloDB[bestEloMatch[0]]
	p2Index := eloDB[bestEloMatch[1]]
	fmt.Println("indexes")
	fmt.Println(p1Index, p2Index)
	p1 := matchMaker.playersPool[p1Index]
	p2 := matchMaker.playersPool[p2Index]

	m := match.NewMatch([2]*player.Player{p1, p2})
	return m, matchMaker
}

func sortMapByElo(elos map[string]int) map[string]int {
	keys := make([]string, 0, len(elos))
	for k := range elos {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return elos[keys[i]] < elos[keys[j]]
	})
	return elos
}

func calculateElo(m *match.Match, k int) {
	m.Winner.Elo += k
	if m.Winner == m.Players[0] {
		m.Players[1].Elo -= k
	} else {
		m.Players[0].Elo -= k
	}
}

func firstRound(mM *MatchMaker) *match.Match {
	fmt.Println("hewe")
	length := len(mM.playersPool)

	fmt.Println(len(mM.playersPool), mM.playersPool)
	n1 := Random(length)
	n2 := Random(length)
	for n2 == n1 {
		n2 = Random(length)
	}
	fmt.Println(n1, n2)
	p1 := mM.playersPool[n1]
	p2 := mM.playersPool[n2]
	if n1 > n2 {
		helpers.DeleteSliceElement(&mM.playersPool, n1)
		helpers.DeleteSliceElement(&mM.playersPool, n2)
	} else {
		helpers.DeleteSliceElement(&mM.playersPool, n2)
		helpers.DeleteSliceElement(&mM.playersPool, n1)
	}

	players := [2]*player.Player{p1, p2}
	m := match.NewMatch(players)
	calculateElo(m, mM.k)
	fmt.Println("pr")
	if len(mM.playersPool) == 0 {
		mM.round = 1
		return m
	}

	return m
}

func Random(n int) int {
	return rand.Intn(n)
}

func findBestEloMatch(elos []int) []int {
	slices.Sort(elos)
	minDiff := math.MaxInt
	// fmt.Println(elos)
	var res [][]int
	for i := range len(elos) - 1 {
		// fmt.Println("here 5")
		diff := elos[i+1] - elos[i]

		// fmt.Println(diff)
		if diff <= minDiff {

			// fmt.Println("here 9")
			minDiff = diff

			tempArr := []int{elos[i], elos[i+1]}
			res = append(res, tempArr)

			// fmt.Println("here 2")
			if minDiff <= (res[0][1] - res[0][0]) {
				// fmt.Println("here 3")
				res = res[:1]
			}

		}

	}
	// fmt.Println(res)
	return res[0]
}
