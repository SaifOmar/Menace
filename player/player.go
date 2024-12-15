package player

import (
	"fmt"
	// "time"
	// "math/rand"
)

type Player struct {
	Hp      int
	Name    string
	Height  float32
	Weight  float32
	Strengh int
	Stamina int
	Iq      int
	Elo     int
	Passive Ability
	Ability Ability
	Record
}

func NewPlayer(name string, height float32, weight float32, strength int, stamina int, iq int, passiveAbility Ability, firstAbility Ability) *Player {
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

func CreatePlayers() []*Player {
	passive := NewAbility("fart", "fart", 2)
	ability := NewAbility("muh", "dah", 10)

	player1 := NewPlayer("Salah", 192, 120, 3, 4, -120, *passive, *ability)
	player2 := NewPlayer("Yousry", 175, 75, 3, 4, -120, *passive, *ability)

	return []*Player{player1, player2}
}

func (player Player) Hit(oponnent *Player) *Player {
	damage := player.Passive.Damage
	fmt.Println(player.Name, "has hit", oponnent.Name, "for", damage)
	oponnent.getHit(damage)
	return oponnent
}

func (player Player) AbilityHit(oponnent *Player) *Player {
	damage := player.Ability.Damage
	fmt.Println(player.Name, "has hit", oponnent.Name, "for", damage)
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
