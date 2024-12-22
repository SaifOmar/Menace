package player

import (
// "fmt"
// "time"
// "math/rand"
)

type Player struct {
	Hp  int
	Elo int

	Strengh int
	Stamina int
	Iq      int
	Height  int
	Weight  int
	Name    string

	Passive Ability
	Ability Ability

	Record
}

func NewPlayer(name string, height int, weight int, strength int, stamina int, iq int, passiveAbility Ability, firstAbility Ability) *Player {
	player := &Player{
		Hp:      100,
		Elo:     900,
		Name:    name,
		Height:  height,
		Weight:  weight,
		Strengh: strength,
		Stamina: stamina,
		Iq:      iq,
		Passive: passiveAbility,
		Ability: firstAbility,
	}
	return player
}

func (player Player) Hit(oponnent *Player) *Player {
	damage := player.Passive.Damage
	// fmt.Println(player.Name, "has hit", oponnent.Name, "for", damage)
	oponnent.getHit(damage)
	return oponnent
}

func (player Player) AbilityHit(oponnent *Player) *Player {
	damage := player.Ability.Damage
	// fmt.Println(player.Name, "has hit", oponnent.Name, "for", damage)
	oponnent.getHit(damage)
	return oponnent
}

func (player *Player) getHit(damage int) int {
	player.Hp -= damage
	if player.Hp < 0 {
		player.Hp = 0
	}
	return player.Hp
}
