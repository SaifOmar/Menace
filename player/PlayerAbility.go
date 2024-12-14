package player

type Ability struct {
	Name   string
	Damage int
	Effect string
}

func (ability *Ability) NewAbility(name string, damage int, effect string) *Ability {
	ability.Name = name
	ability.Damage = damage
	ability.Effect = effect
	return ability
}
