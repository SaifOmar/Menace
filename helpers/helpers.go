package helpers

import (
	"math/rand"
	// "fmt"

	"TournamentProject/player"
)

func DeleteSliceElement(slice *[]*player.Player, index int) *[]*player.Player {
	newSlice := *slice
	if len(newSlice) == 1 {
		*slice = make([]*player.Player, 0)

		return slice
	}
	newSlice = append(newSlice[:index], newSlice[index+1:]...)

	*slice = newSlice
	return slice
}

func Random(n int) int {
	if n == 1 {
		n++
	}
	return rand.Intn(n)
}
