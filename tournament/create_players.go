package tournament

import (
	// "math/rand"

	"TournamentProject/helpers"
	"TournamentProject/player"
)

// the values for these should come from some sort of database or config file
var playerData = map[int][]interface{}{
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

	0: {"slta", 185, 85, 70, 40, 120},
	1: {"salah", 192, 120, 60, 40, 114},
	2: {"yousry", 172, 65, 50, 60, 40},
	3: {"omar", 172, 75, 30, 40, 140},
	// 1.name 2.height 3.weight 4.strenght 5.stamina 6.iq
	4: {"joe", 180, 130, 60, 10, 40},
	5: {"mahmod", 183, 70, 50, 85, 120},
	6: {"mo7ie", 182, 90, 50, 22, 45},
	7: {"fr5a", 171, 60, 30, 100, 60},
}

var abilitiesData = map[string][]*player.Ability{
	"salah": {player.NewAbility("Not Even Close", "Salah is a shitter", 1), player.NewAbility("Calculated", "Skill is a myth it's all luck", helpers.RandomSecondAbilityDmg(0, 40))},

	"yousry": {player.NewAbility("Ysta m3ah flash", "Yosry Has healed you", 5), player.NewAbility("Darrag el agnab", "Ignite top lane", helpers.RandomSecondAbilityDmg(10, 30))},

	"joe": {player.NewAbility("Brain Dmg", "Joe has donated brain dmg to you", 9), player.NewAbility("Lock Screen", "Joe is watching someone else's screen", helpers.RandomSecondAbilityDmg(15, 20))},

	"omar": {player.NewAbility("Bad Dad Joke BOMB", "Say Goodbye to your balls", 10), player.NewAbility("Retardness", "Omar has caused retardedness", helpers.RandomSecondAbilityDmg(15, 20))},

	"slta": {player.NewAbility("El dagger e5tfa", "Fe tez men", 10), player.NewAbility("Ya 3aaaaaaaaaaaaaaaam", "Report far5a", helpers.RandomSecondAbilityDmg(10, 20))},

	"mahmod": {player.NewAbility("IQ LOSS", "7oda has traded his IQ for more dmg, (d5l bdma8o)", 10), player.NewAbility("Spotify", "Ya sa7by bt5onk kman bt2oly t3ala 3yzak", helpers.RandomSecondAbilityDmg(15, 20))},

	"fr5a": {player.NewAbility("EGGNITE", "Far5a laid an egg", 8), player.NewAbility("Wakaak", "Wak wak wakaaaaak", helpers.RandomSecondAbilityDmg(15, 25))},

	"mo7ie": {player.NewAbility("SubHuman", "What is this thing", 3), player.NewAbility("Absence", "Is anyone there?", helpers.RandomSecondAbilityDmg(1, 33))},
}

func CreatePlayers() []*player.Player {
	length := len(playerData)

	var players []*player.Player

	for k := range length {
		playerArr, exists := playerData[k]

		if !exists {
			return nil
		}
		// this entire block is just because go sucks
		n, nOk := playerArr[0].(string)
		h, hOk := playerArr[1].(int)
		w, wOk := playerArr[2].(int)
		s, sOk := playerArr[3].(int)
		st, stOk := playerArr[4].(int)
		iq, k := playerArr[5].(int)
		if !(nOk && hOk && wOk && sOk && stOk && k) {
			return nil
		}

		ability, passive := createAbilities(n)
		player := player.NewPlayer(n, h, w, s, st, iq, *ability, *passive)
		players = append(players, player)

	}

	return players
}

func createAbilities(n string) (*player.Ability, *player.Ability) {
	playerAbilities, exists := abilitiesData[n]

	if !exists {
		panic("Player abilities not found")
		// return nil, nil
	}

	return playerAbilities[0], playerAbilities[1]
}
