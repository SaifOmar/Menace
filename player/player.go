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

	SkillLevel float64

	// Record
}

// !FIX very very very bad should be a builder pattern instead of this mess

func NewPlayer(name string) *Player {
	player := &Player{
		Name:        name,
		Hp:          100, // calculated attr
		Elo:         900, // base stat(something to calc dmg)
		AdjustedElo: 900,
		WP:          100.00,
	}
	player.SkillLevel = player.calculateSkillLevel()
	return player
}

func (p *Player) AddHeight(n int) *Player {
	p.Height = n
	return p
}

func (p *Player) AddWeight(n int) *Player {
	p.Weight = n
	return p
}

func (p *Player) AddStamina(n int) *Player {
	p.Stamina = n
	return p
}

func (p *Player) AddStrength(n int) *Player {
	p.Strength = n
	return p
}

func (p *Player) AddIq(n int) *Player {
	p.Iq = n
	return p
}

func (p *Player) AttatchAbilites(a1 Ability, a2 Ability) *Player {
	p.Ability = a1
	p.Passive = a2
	return p
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
	// calculate the skill level based on player normalized stat values
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
	// normalize values
	for k := range v {
		// if the value is more than the max it sets it to the max
		if v[k] > max[k] {
			v[k] = max[k]
		}
		// if the value is less than the min it sets it to the min
		if v[k] < min[k] {
			v[k] = min[k]
		}
		// normalize the value
		v[k] = (v[k] - min[k]) / (max[k] - min[k])
	}

	return v
}

// Arrmor , Dmg
// hit (200 %80 less dmg)
// ab1 dmg + ab2 dmg
