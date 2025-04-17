package luck

type EffectResult struct {
	Damage int
	Note   string
	Player int // 0 for player 1, 1 for player 2
}

func NewEffectResult(player int, note string) *EffectResult {
	return &EffectResult{
		Damage: 0,
		Note:   note,
		Player: player,
	}
}

func (e *EffectResult) SetDamage(damage int) {
	e.Damage = damage
}
