package tournament

import (
	"fmt"

	"TournamentProject/db"
	"TournamentProject/helpers"
	"TournamentProject/match"
	"TournamentProject/player"
	"TournamentProject/ranks"
)

type Tournament struct {
	*db.Db
	Matches    []match.Match
	s          string
	Players    []*player.Player
	Winner     *player.Player
	Logger     *helpers.TournamentLogger
	matchMaker *MatchMaker
	// Records    []player.Record
}

func NewTournament(logger *helpers.TournamentLogger) *Tournament {
	t := &Tournament{
		Logger: logger,
		s:      "DEFAULT",
	}
	return t
}

func (t *Tournament) UseInMem() {
	t.s = "MEM"
}

func (t *Tournament) Run() *Tournament {
	t.addPlayers().start()
	return t
}

func (t *Tournament) addPlayers() *Tournament {
	switch t.s {
	case "DEFAULT":
		t.Players = CreatePlayers()
	case "MEM":
		t.Players = t.InMemPlayers()
	}
	return t
}

func (t *Tournament) start() *Tournament {
	t.Logger.Debug("StartDebugging")
	t.matchMaker = NewMatchMaker(t.Players, t.Logger)
	for {
		m := t.matchMaker.MakeMatch()
		t.saveMatch(m)

		if m.Winner.Elo >= ranks.Platinum {
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

// how this works will be completely changed I believe
func (t *Tournament) saveMatch(m *match.Match) {
	t.Logger.Debug("Saving Match")
	t.Matches = append(t.Matches, *m)
	// records := generatePlayerRecords(m)
	// t.Records = append(t.Records, records...)
}

// func generatePlayerRecords(m *match.Match) []player.Record {
// 	var records []player.Record
// 	for _, p := range m.Players {
// 		winner := (m.Winner.Name == p.Name)
// 		p.RecordArr = append(p.RecordArr, winner)
// 		records = append(records, p.Record)
// 		// p.CalculateWinPercentage()
// 		p.EvaluatePlayerPerformance(0)
// 	}
// 	return records
// }
