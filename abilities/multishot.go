package abilities

import (
	"huntsim/player"
	"huntsim/util"
)

type MultiShot struct {
	BaseCastTime float64
}

func (ms MultiShot) Damage(p *player.Player) float64 {
	// Is it a miss?
	if util.RollDice(p.MissChance()) {
		return 0
	}

	// Calculate Base Damage
	weapon := p.Equipment.Ranged
	unmodifiedDamage := weapon.AverageDamage() + weapon.AmmoDamage() + p.EffectiveAP()*0.2 + float64(weapon.Scope.Damage) + 205

	// Is it a crit?
	if util.RollDice(p.CritChance()) {
		unmodifiedDamage = unmodifiedDamage * p.RangeCritDamageModifier()
	}

	return unmodifiedDamage * FOCUSED_FIRE_MOD
}

func NewMultiShot() *MultiShot {
	return &MultiShot{
		BaseCastTime: 0.5,
	}
}
