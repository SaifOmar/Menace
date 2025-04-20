package player

type Ability struct {
	Name   string `json:"name"`
	Effect string `json:"effect"`
	Damage int    `json:"damage"`
}

func NewAbility(name string) *Ability {
	return &Ability{
		Name: name,
	}
}

func (a *Ability) AddDamage(damage int) *Ability {
	a.Damage = damage
	return a
}

func (a *Ability) AddEffect(effect string) *Ability {
	a.Effect = effect
	return a
}
