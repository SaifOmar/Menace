package tournament

import (
	"fmt"
	// "runtime"
	// "strings"

	"TournamentProject/match"
	"TournamentProject/player"
)

func CreatePlayers() []*player.Player {
	var players []*player.Player
	playerData := map[int][]interface{}{
		0: {"slta", 185, 85, 10, 10, 200},
		1: {"salah", 192, 120, 3, 4, -120},
		2: {"yousry", 192, 120, 3, 4, -120},
		3: {"okl", 192, 120, 3, 4, -120},

		4: {"joe", 192, 120, 3, 4, -120},
		5: {"mahmod", 192, 120, 3, 4, -120},
		6: {"mo7ie", 192, 120, 3, 4, -120},
		7: {"fr5a", 192, 120, 3, 4, -120},
	}

	length := len(playerData)
	for k := range length {

		playerArr, exists := playerData[k]

		if !exists {
			return nil
		}

		n, nOk := playerArr[0].(string)
		h, hOk := playerArr[1].(int)
		w, wOk := playerArr[2].(int)
		s, sOk := playerArr[3].(int)
		st, stOk := playerArr[4].(int)
		iq, k := playerArr[5].(int)

		if !(nOk && hOk && wOk && sOk && stOk && k) {
			return nil
		}
		a, p := createAbilities(n)
		pl := player.NewPlayer(n, h, w, s, st, iq, *a, *p)
		players = append(players, pl)

	}

	return players
}

func createAbilities(n string) (*player.Ability, *player.Ability) {
	abilitiesData := map[string][]*player.Ability{
		"salah":  {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"yousry": {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"joe":    {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"okl":    {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"slta":   {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"mahmod": {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"fr5a":   {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
		"mo7ie":  {player.NewAbility("muh", "du", 10), player.NewAbility("some", "fu", 19)},
	}

	playerAbilities, exists := abilitiesData[n]

	if !exists {
		return nil, nil
	}

	return playerAbilities[0], playerAbilities[1]
}

type Tournament struct {
	Matches []match.Match
	Records []player.Record
	Players []*player.Player
	mm      *MatchMaker
	Winner  *player.Player
}

func NewTournament() *Tournament {
	Players := CreatePlayers()
	t := &Tournament{
		Players: Players,
		mm:      NewMatchMaker(Players),
	}

	for range 1000 {

		m, _ := t.mm.MakeMatch()
		t.saveMatch(m)

		if m.Winner.Elo >= 3200 {
			t.Winner = m.Winner
			fmt.Println(t.Winner.Name, "has won and his elo is: ", t.Winner.Elo, "\n")
			for _, p := range t.Players {
				fmt.Println(p.Name, ": ", p.Elo)
			}
			break
		}
	}

	return t
}

func (t *Tournament) saveMatch(m *match.Match) {
	t.Matches = append(t.Matches, *m)
	records := generatePlayerRecords(m)
	t.Records = append(t.Records, records...)
}

func generatePlayerRecords(m *match.Match) []player.Record {
	var records []player.Record
	for _, p := range m.Players {
		winner := (m.Winner.Name == p.Name)
		p.RecordArr = append(p.RecordArr, winner)
		records = append(records, p.Record)
		// p.CalculateWinPercentage()
		p.EvaluatePlayerPerformance(0)

	}

	return records
}
