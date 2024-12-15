package tournament

import (
	// "fmt"

	"TournamentProject/match"
	"TournamentProject/player"
)

type Tournament struct {
	Matches []match.Match
	Records []player.Record
	Players []*player.Player
	Winner  player.Player
}

func NewTournament(players []*player.Player) (*Tournament, *match.Match) {
	tM := &Tournament{
		Players: players,
	}
	_, match := NewMatchMaker(tM.Players)
	if match.Winner.Name != "" {
		tM.Winner = match.Winner
	}
	return tM, match
}
