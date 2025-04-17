package helpers

import (
	"math"
	"math/rand"
	"slices"

	"TournamentProject/player"
)

func DeleteSliceElement(slice *[]*player.Player, index int) *[]*player.Player {
	newSlice := *slice
	if len(newSlice) == 1 {
		*slice = make([]*player.Player, 0)

		return slice
	}
	newSlice = slices.Delete(newSlice, index, index+1)

	*slice = newSlice
	return slice
}

func Random(n int) int {
	if n == 1 {
		n++
	}
	return rand.Intn(n)
}

// sorts players based on Adjusted Elo
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

func minDiffThatShit(playerslice *[]*player.Player) []*player.Player {
	s := *playerslice
	l := len(s)

	QuickSort(playerslice, 0, l-1)

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

func FindCandidates(players []*player.Player) (*player.Player, *player.Player) {
	// this will find the two players with the lowest Elo diff
	// and return them
	lowestDiff := math.MaxInt
	var player1, player2 *player.Player

	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			diff := int(math.Abs(float64(players[i].Elo - players[j].Elo)))
			if diff < lowestDiff {
				lowestDiff = diff
				player1 = players[i]
				player2 = players[j]
			}
		}
	}

	return player1, player2
}

func RandomSecondAbilityDmg(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
