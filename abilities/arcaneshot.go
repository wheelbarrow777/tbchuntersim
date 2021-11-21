package abilities

import (
	"tbchuntersim/player"
	"tbchuntersim/util"

	log "github.com/sirupsen/logrus"
)

const (
	arcaneShotBonus = 273
)

type ArcaneShot struct {
	BaseCastTime    float64
	CurrentCooldown float64
}

func (as ArcaneShot) Cast(p *player.Player) float64 {
	log.Debug("Casting Arcane Shot")

	// Is it a miss?
	if util.RollDice(p.MissChance()) {
		return 0
	}

	unmodifiedDamage := p.EffectiveAP()*0.15 + arcaneShotBonus

	// TODO: Curse of Shadow Modified should not always be applied
	unmodifiedDamage = unmodifiedDamage * CURSE_OF_SHADOW_MOD * MISERY_MOD

	// Is it a crit?
	if util.RollDice(p.CritChance()) {
		unmodifiedDamage = unmodifiedDamage * p.RangeCritDamageModifier()
	}

	return unmodifiedDamage * FOCUSED_FIRE_MOD
}

func NewArcaneShot() *ArcaneShot {
	return &ArcaneShot{
		BaseCastTime: 0.5,
	}
}
