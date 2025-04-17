package tournament

import (
	"fmt"

	"TournamentProject/helpers"
	"TournamentProject/match"
	"TournamentProject/player"
	"TournamentProject/ranks"
)

// func validatePlayerAttributes() {
// }

type Tournament struct {
	Matches    []match.Match
	Records    []player.Record
	Players    []*player.Player
	Winner     *player.Player
	Logger     *helpers.TournamentLogger
	matchMaker *MatchMaker
}

func NewTournament(logger *helpers.TournamentLogger) *Tournament {
	t := &Tournament{
		Logger: logger,
	}
	return t.start()
}

func (t *Tournament) start() *Tournament {
	t.Logger.Debug("StartDebugging")
	t.matchMaker = NewMatchMaker(CreatePlayers(), t.Logger)
	for {
		// this will make the match and start it and only return to save only when It's finished
		// should I change this to be able to make the match and start it after from here
		m := t.matchMaker.MakeMatch()
		t.saveMatch(m)

		if m.Winner.Elo >= ranks.Challenger {
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
