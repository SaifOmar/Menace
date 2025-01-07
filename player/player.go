package player

type Player struct {
	Name string

	Hp  int
	Elo int

	AdjustedElo int

	Strength int
	Stamina  int
	Iq       int
	Height   int
	Weight   int

	Passive Ability
	Ability Ability

	NMatches int
	WP       float64
	WinCount int

	skillLevel float64

	Record
}

func NewPlayer(name string, height int, weight int, strength int, stamina int, iq int, passiveAbility Ability, firstAbility Ability) *Player {
	player := &Player{
		Hp:          100, // calculated attr
		Elo:         900, // base stat(something to calc dmg)
		Name:        name,
		Height:      height,
		Weight:      weight,
		Strength:    strength,
		Stamina:     stamina,
		Iq:          iq,
		Passive:     passiveAbility,
		Ability:     firstAbility,
		AdjustedElo: 900,
		WP:          100.00,
	}
	player.skillLevel = player.calculateSkillLevel()
	return player
}

func (player Player) Hit(oponnent *Player) *Player {
	damage := player.Passive.Damage
	oponnent.getHit(damage)
	return oponnent
}

func (player Player) AbilityHit(oponnent *Player) *Player {
	damage := player.Ability.Damage
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

func (p Player) calculateSkillLevel() float64 {
	v := map[string]float64{
		"iq":       float64(p.Iq),
		"stamina":  float64(p.Stamina),
		"strength": float64(p.Strength),
		"weight":   float64(p.Weight),
		"height":   float64(p.Height),
	}
	w := map[string]float64{
		"iq":       0.3,
		"stamina":  0.15,
		"strength": 0.3,
		"weight":   0.125,
		"height":   0.125,
	}

	v = normlizeValues(v)

	skillLevel := 0.0
	for k := range v {
		skillLevel = skillLevel + (w[k] * float64(v[k]))
	}
	return skillLevel
}

func normlizeValues(v map[string]float64) map[string]float64 {
	min := map[string]float64{
		"iq":       40,
		"stamina":  10,
		"strength": 10,
		"height":   135,
		"weight":   50,
	}
	max := map[string]float64{
		"iq":       180,
		"stamina":  100,
		"strength": 100,
		"height":   222,
		"weight":   140,
	}
	for k := range v {
		if v[k] > max[k] {
			v[k] = max[k]
		}
		if v[k] < min[k] {
			v[k] = min[k]
		}
		v[k] = (v[k] - min[k]) / (max[k] - min[k])
	}

	return v
}

// Arrmor , Dmg
// hit (200 %80 less dmg)
// ab1 dmg + ab2 dmg
