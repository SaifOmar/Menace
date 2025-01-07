package tournament

import (
	"fmt"

	"TournamentProject/helpers"
	"TournamentProject/match"
	"TournamentProject/player"
)

func validatePlayerAttributes() {
}

func CreatePlayers() []*player.Player {
	var players []*player.Player

	playerData := map[int][]interface{}{
		// 1.name 2.height 3.weight 4.strenght 5.stamina 6.iq

		// min := map[string]float64{
		// 	"iq":       40,
		// 	"stamina":  10,
		// 	"strength": 10,
		// 	"height":   135,
		// 	"weight":   50,
		// }
		// max := map[string]float64{
		// 	"iq":       180,
		// 	"stamina":  100,
		// 	"strength": 100,
		// 	"height":   222,
		// 	"weight":   140,
		// }
		0: {"slta", 222, 85, 100, 100, 180},
		1: {"salah", 192, 120, 50, 30, 114},
		2: {"yousry", 172, 65, 40, 10, 40},
		3: {"omar", 172, 75, 33, 40, 140},

		4: {"joe", 180, 130, 11, 10, 90},
		5: {"mahmod", 183, 70, 65, 85, 120},
		6: {"mo7ie", 182, 90, 30, 22, 45},
		7: {"fr5a", 171, 60, 10, 60, 41},
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
		"salah":  {player.NewAbility("Not Even Close", "Salah is a shitter", 2), player.NewAbility("Calculated", "Luck is a myth it's all skill", 1)},
		"yousry": {player.NewAbility("Ysta m3ah flash", "Yosry Has healed you", 0), player.NewAbility("Darrag el agnab", "Ignite top lane", 30)},
		"joe":    {player.NewAbility("Brain Dmg", "Joe has donated brain dmg to you", 8), player.NewAbility("Lock Screen", "Joe is watching someone else's screen", 12)},
		"omar":   {player.NewAbility("Bad Dad Joke BOMB", "Say Goodbye to your balls", 8), player.NewAbility("Retardness", "Omar has caused retardedness", 10)},
		"slta":   {player.NewAbility("El dagger e5tfa", "Fe tez men", 10), player.NewAbility("Ya 3aaaaaaaaaaaaaaaam", "Report far5a", 19)},
		"mahmod": {player.NewAbility("IQ LOSS", "7oda has traded his IQ for more dmg, (d5l bdma8o)", 40), player.NewAbility("Spotify", "Ya sa7by bt5onk kman bt2oly t3ala 3yzak", 20)},
		"fr5a":   {player.NewAbility("EGGNITE", "Far5a laid an egg", 11), player.NewAbility("Wakaak", "Wak wak wakaaaaak", 9)},
		"mo7ie":  {player.NewAbility("SubHuman", "What is this thing", 10), player.NewAbility("Absence", "Is anyone there?", 11)},
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
	Logger  *helpers.TournamentLogger
}

func NewTournament(logger *helpers.TournamentLogger) *Tournament {
	Players := CreatePlayers()
	t := &Tournament{
		Players: Players,
		Logger:  logger,
		mm:      NewMatchMaker(Players, logger),
	}
	// t.Logger.Debug("StartDebugging")

	for {
		m, _ := t.mm.MakeMatch()
		t.saveMatch(m)

		if m.Winner.Elo >= 3200 {
			t.Winner = m.Winner
			t.Logger.Info(fmt.Sprintf("%s Has won the tournament and with elo of: %d ", t.Winner.Name, t.Winner.Elo))
			t.Logger.Info("Players : 👇")
			for _, p := range t.Players {
				t.Logger.Info(fmt.Sprintf("Player: %s | Elo : %d | After : %d", p.Name, p.Elo, p.AdjustedElo))
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
