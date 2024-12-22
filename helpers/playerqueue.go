package helpers

import (
	"fmt"

	"TournamentProject/player"
)

type PlayerQueue []*player.Player

func (q *PlayerQueue) Enqueue(value *player.Player) {
	*q = append(*q, value)
}

func (q *PlayerQueue) Dequeue() (*player.Player, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("player queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}
