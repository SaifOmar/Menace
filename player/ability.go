package player

type Ability struct {
	Name   string
	Effect string
	Damage int
}

func NewAbility(name string, effect string, damage int) *Ability {
	return &Ability{
		Name:   name,
		Effect: effect,
		Damage: damage,
	}
}
