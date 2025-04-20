package tournament

import (
	"fmt"
	"strconv"
	"strings"

	"TournamentProject/helpers"
	"TournamentProject/player"
)

// the values for these should come from some sort of database or config file
var playerData = []string{
	"slta",
	"salah",
	"yousry",
	"omar",
	"joe",
	"mahmod",
	"mo7ie",
	"fr5a",
}

func CreatePlayers() []*player.Player {
	var players []*player.Player
	for _, p := range playerData {
		fmt.Println(p)
		newPlayer := player.NewPlayer(p).
			AddStamina(Players[p]["satmina"]).
			AddStrength(Players[p]["strength"]).
			AddWeight(Players[p]["weight"]).
			AddIq(Players[p]["iq"]).
			AddHeight(Players[p]["height"]).
			AttatchAbilites(createAbilities(p))
		fmt.Println(newPlayer)
		players = append(players, newPlayer)
	}

	return players
}

func createAbilities(p string) (player.Ability, player.Ability) {
	abilities, exists := PlayerAbilities[p]
	if !exists {
		panic("Can't have a player that doesnt have abilites")
	}
	var abilitiesSlice []player.Ability
	for _, a := range abilities {
		ability := player.NewAbility(a["name"]).
			AddEffect(a["effect"])
		var dmg int
		dmgString := a["damage"]
		if strings.Contains(dmgString, "-") {
			parts := strings.Split(dmgString, "-")
			if len(parts) == 2 {
				randStart, err := strconv.Atoi(parts[0])
				if err == nil {
					randEnd, err := strconv.Atoi(parts[1])
					if err == nil {
						dmg = helpers.RandomSecondAbilityDmg(randStart, randEnd)
					}
				}
			}
		} else {
			if temp, err := strconv.Atoi(dmgString); err != nil {
				panic(err)
			} else {
				dmg = temp
			}
		}
		ability.AddDamage(dmg)
		abilitiesSlice = append(abilitiesSlice, *ability)
	}
	// fmt.Println(abilitiesSlice)
	return abilitiesSlice[0], abilitiesSlice[1]
}
