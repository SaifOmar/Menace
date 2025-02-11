package match

import (
	"fmt"
	"math/rand"

	"TournamentProject/helpers"
	"TournamentProject/player"
)

type Luck struct{}

type Luckk interface {
	// very very very low chance salah()
	OneShot() int // one shots the player (1k dmg or sm like that)
	JumpOff() int // one shots the player (1k dmg or sm like that)

	// low chance jackBot()
	DonateBrainDmg() string //  prinst you have been given the gift of brain dmg and -20 hp
	Reflect() int           // reflects the dmg
	Immunity() int          // player takes no dmg

	// normal luckyBitch()
	OutPlay() string     // p1 has flashed(message)
	Heal() (int, string) // 3rbyt es3af
	Armor() int          // aka malphite
	Crit() int
}

func NewLuck() *Luck {
	l := Luck{}
	return &l
}

func (l *Luck) doSm(p1Sk, p2Sk float64) int {
	return 0
}

func OneShot(p *player.Player) {
}

func Die(p *player.Player) {
}

func Immunity(p *player.Player) {
}

func Reflect(p *player.Player) {
}

func Armor(p *player.Player) {
}

func Heal(p *player.Player) {
}

func Crit(p *player.Player) {
}

func OutPlay(p *player.Player) {
}

func DonateBrainDmg() {
}

func getLucky() {
	n := helpers.Random(100)
	switch {

	// 1%
	case n == 50:
		{
			salah()
			break
		}
	// 20%
	case n < 20:
		{
			jackBot()
			break
		}

	// 40%
	case n <= 40:
		{
			luckyBitch()
			break
		}
	}

	// if a player gets lucky he can call to one of the functions in the interface
	// to be able to do determine which one of the functions I can call I need to have a map in my head for that which for now I don't have
	// highest out come is a player one shoting, or jumping off (committing suicide)
	// after that comes brain dmg, reflect, immunity
	// normal cases are just heal, armor, crit and outplaying
	// I call the base function here
	if n <= 0.8 {
		fmt.Print("hello mom")
	}
}

func salah() {
	if helpers.Random(2) == 1 {
		OneShot()
	} else {
		Die()
	}
}

func jackBot() {
	switch helpers.Random(3) {
	case 0:
		{
			DonateBrainDmg()
			break
		}
	case 1:
		{
			Reflect()

			break
		}
	case 2:
		{
			Immunity()
			break
		}
	}
}

func luckyBitch() {
	switch helpers.Random(4) {
	case 0:
		{
			Outplay()
			break
		}
	case 1:
		{
			Heal()

			break
		}
	case 2:
		{
			Armor()
			break
		}
	case 3:
		{
			Crit()
			break
		}
	}
}
