package helpers

import (
	"fmt"
	// "fmt"

	"TournamentProject/player"
)

func DeleteSliceElement(slice *[]*player.Player, index int) *[]*player.Player {
	newSlice := *slice
	if len(newSlice) == 1 {
		*slice = make([]*player.Player, 0)
		fmt.Println("here")

		return slice
	}
	newSlice = append(newSlice[:index], newSlice[index+1:]...)
	*slice = newSlice
	return slice
}
