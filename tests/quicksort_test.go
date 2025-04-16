package tournament_test

import (
	"testing"

	"TournamentProject/tournament"
)

func TestQuickSort(t *testing.T) {
	players := tournament.CreatePlayers()
	low := 0
	high := 0
	tournament.QuickSort(&players, low, high)

	for i := range len(players) - 1 {
		if players[i].AdjustedElo < players[i+1].AdjustedElo {
			t.Errorf("Players not sorted correctly: %d < %d at index %d", players[i].AdjustedElo, players[i+1].AdjustedElo, i)
		}
	}
}
