package player

import (
	"fmt"
	"math/rand"
)

// import "fmt"

type Player struct {
	Hp      int
	Name    string
	Height  float32
	Weight  float32
	Strengh int
	Stamina int
	Iq      int
	Passive Ability
	Ability Ability
}

func (player *Player) NewPlayer(name string, height float32, weight float32, strength int, stamina int, iq int, passiveAbility Ability, firstAbility Ability) *Player {
	player.Hp = 100
	player.Name = name
	player.Height = height
	player.Weight = weight
	player.Strengh = strength
	player.Stamina = stamina
	player.Iq = iq
	player.Passive = passiveAbility
	player.Ability = firstAbility
	return player
}

func (player Player) Hit(oponnent *Player) *Player {
	damage := rand.Intn(16)
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
