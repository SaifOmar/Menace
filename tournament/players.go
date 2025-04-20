package tournament

//	min := map[string]float64{
//		"iq":       40,
//		"stamina":  10,
//		"strength": 10,
//		"height":   135,
//		"weight":   50,
//	}
//
//	max := map[string]float64{
//		"iq":       180,
//		"stamina":  100,
//		"strength": 100,
//		"height":   222,
//		"weight":   140,
//	}
//
// 1. height 2. weight 3. strength 4. stamina 5. iq
var Players = map[string]map[string]int{
	"slta": {
		"height":   185,
		"weight":   85,
		"strength": 70,
		"stamina":  40,
		"iq":       120,
	},
	"salah": {
		"height":   192,
		"weight":   120,
		"strength": 60,
		"stamina":  40,
		"iq":       114,
	},
	"yousry": {
		"height":   172,
		"weight":   65,
		"strength": 50,
		"stamina":  60,
		"iq":       40,
	},
	"omar": {
		"height":   172,
		"weight":   75,
		"strength": 30,
		"stamina":  40,
		"iq":       140,
	},
	"joe": {
		"height":   180,
		"weight":   130,
		"strength": 60,
		"stamina":  10,
		"iq":       40,
	},
	"mahmod": {
		"height":   183,
		"weight":   70,
		"strength": 50,
		"stamina":  85,
		"iq":       120,
	},
	"mo7ie": {
		"height":   182,
		"weight":   90,
		"strength": 50,
		"stamina":  22,
		"iq":       45,
	},
	"fr5a": {
		"height":   171,
		"weight":   60,
		"strength": 30,
		"stamina":  100,
		"iq":       60,
	},
}

// returns an array of the player abilites maps
var PlayerAbilities = map[string][]map[string]string{
	"salah": {
		{"name": "Not Even Close", "effect": "Salah is a shitter", "damage": "1"},
		{"name": "Calculated", "effect": "Skill is a myth it's all luck", "damage": "0-40"},
	},
	"yousry": {
		{"name": "Ysta m3ah flash", "effect": "Yosry Has healed you", "damage": "5"},
		{"name": "Darrag el agnab", "effect": "Ignite top lane", "damage": "10-30"},
	},
	"joe": {
		{"name": "Brain Dmg", "effect": "Joe has donated brain dmg to you", "damage": "9"},
		{"name": "Lock Screen", "effect": "Joe is watching someone else's screen", "damage": "15-20"},
	},
	"omar": {
		{"name": "Bad Dad Joke BOMB", "effect": "Say Goodbye to your balls", "damage": "10"},
		{"name": "Retardness", "effect": "Omar has caused retardedness", "damage": "15-20"},
	},
	"slta": {
		{"name": "El dagger e5tfa", "effect": "Fe tez men", "damage": "10"},
		{"name": "Ya 3aaaaaaaaaaaaaaaam", "effect": "Report far5a", "damage": "10-20"},
	},
	"mahmod": {
		{"name": "IQ LOSS", "effect": "7oda has traded his IQ for more dmg, (d5l bdma8o)", "damage": "10"},
		{"name": "Spotify", "effect": "Ya sa7by bt5onk kman bt2oly t3ala 3yzak", "damage": "15-20"},
	},
	"fr5a": {
		{"name": "EGGNITE", "effect": "Far5a laid an egg", "damage": "8"},
		{"name": "Wakaak", "effect": "Wak wak wakaaaaak", "damage": "15-25"},
	},
	"mo7ie": {
		{"name": "SubHuman", "effect": "What is this thing", "damage": "3"},
		{"name": "Absence", "effect": "Is anyone there?", "damage": "1-33"},
	},
}
