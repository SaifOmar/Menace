package tournament_test

import (
	"reflect"
	"testing"

	"TournamentProject/tournament"
)

func TestCreatePlayers(t *testing.T) {
	players := tournament.CreatePlayers()
	valueType := "saif"
	for _, p := range players {
		if reflect.TypeOf(p.Name) != reflect.TypeOf(valueType) {
			t.Error("Player Name Is not of type string")
		}
	}
}
