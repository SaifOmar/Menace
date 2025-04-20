package luck

//
// //
import (
	"fmt"

	"TournamentProject/helpers"
)

const (
	Legendary = 100
	Epic      = 50
	Rare      = 20
)

type Luck struct {
	Pool             float32
	regenerationPace float32
	Value            string
	EffectResult
}
type Luckk interface {
	// 	// very very very low chance salah() 0.1
	OneShot() int // one shots the player (1k dmg or sm like that)
	JumpOff() int // one shots the player (1k dmg or sm like that)
	//
	// low chance jackBot()
	DonateBrainDmg() string //  prinst you have been given the gift of brain dmg and -20 hp
	Reflect() int           // reflects the dmg
	Immunity() int          // player takes no dmg
	//
	// normal luckyBitch()
	OutPlay() string      // p1 has flashed(message)
	Heal() (int, string)  // 3rbyt es3af
	Armor() (int, string) // malph
	Crit() int            // mulitblies dmg
}

func NewLuck() *Luck {
	return &Luck{
		Pool:             100.0,
		regenerationPace: 20,
		Value:            "Unlucky",
	}
}

func (l *Luck) Regenerate() {
	l.Pool += l.regenerationPace
	if l.Pool > 100 {
		l.Pool = 100
	}
}

func (l *Luck) GetLucky() *Luck {
	if !(l.Pool > 0.0) {
		return l
	}
	n := helpers.Random(100)
	switch {
	// 1%
	case n > 100:
		{
			l.Pool -= Legendary
			l.Value = "Lucky"
			l.salah()
			break
		}

		// 5%
	case (30 <= n) && (n <= 60):
		{
			l.Pool -= Epic
			l.Value = "Lucky"
			l.jackBot()
			break
		}
		// 10%
	case (20 <= n) && (n < 30):
		{
			l.Pool -= Legendary
			l.Value = "Lucky"
			l.luckyBitch()
			break
		}
	}
	return l
}

func (l *Luck) OneShot() {
	player := rollPlayer()
	l.EffectResult = *NewEffectResult(player, "Deleted")
	l.SetDamage(1000)
}

func rollPlayer() int {
	return helpers.Random(2)
}

func (l *Luck) Die() {
	player := rollPlayer()
	l.EffectResult = *NewEffectResult(player, "Bro He just Jumped Off!")
	l.SetDamage(1000)
}

func (l *Luck) Heal() {
	n := helpers.Random(100)
	player := rollPlayer()
	l.EffectResult = *NewEffectResult(player, fmt.Sprintf("healed by %d", n))
	l.SetDamage(-n)
}

func (l *Luck) OutPlay() {
	player := rollPlayer()
	l.EffectResult = *NewEffectResult(player, "Ysta m3ah flash!")
}

func (l *Luck) DonateBrainDmg() {
	player := rollPlayer()
	l.EffectResult = *NewEffectResult(player, "You have been given the gift of brain dmg")
	l.SetDamage(20)
}

func (l *Luck) salah() {
	if helpers.Random(2) == 1 {
		l.OneShot()
	} else {
		l.Die()
	}
}

func (l *Luck) jackBot() {
	switch helpers.Random(1) {
	case 0:
		{
			l.DonateBrainDmg()
			break
		}
	case 1:
		{
			l.DonateBrainDmg()
			break
		}
		// case 1:
		// 	{
		// 		return Reflect()
		//
		// 		break
		// 	}
		// case 2:
		// 	{
		// 		return Immunity()
		// 		break
		// 	}
	}
}

func (l *Luck) luckyBitch() {
	switch helpers.Random(2) {
	case 0:
		{
			l.OutPlay()
			break
		}
	case 1:
		{
			l.Heal()
			break
		}
		// case 2:
		// 	{
		// 		return Armor()
		// 		break
		// 	}
		// case 3:
		// 	{
		// 		return Crit()
		// 		break
		// 	}
	}
}

// func (l *Luck) Immunity() {
// 	l.EffectResult = *NewEffectResult(-100, "Bro He just Jumped Off!")
// 	return
// }

// func Reflect(p *player.Player) {
// }

// func Armor(p *player.Player) {
// }

// func Crit(p *player.Player) {
// }
