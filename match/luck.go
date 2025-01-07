package match

type Luck struct{}

type Luckk interface {
	OneShot() int        // one shots the player (1k dmg or sm like that)
	Heal() (int, string) // 3rbyt es3af
	Armor() int          // aka malphite
	Crit() int
	DonateBrainDmg() string //  prinst you have been given the gift of brain dmg and -20 hp
	Immunity() int          // player takes no dmg
	OutPlay() string        // p1 has flashed(message)
	Reflect() int           // reflects the dmg
}

func NewLuck() *Luck {
	l := Luck{}
	return &l
}

func (l *Luck) doSm(p1Sk, p2Sk float64) int {
	return 0
}

func OneShot() int {
	return 3
}
