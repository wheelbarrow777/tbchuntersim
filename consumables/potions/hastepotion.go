package potions

import (
	"huntsim/abilities"
	"huntsim/player"
	"math"

	log "github.com/sirupsen/logrus"
)

type HastePotion struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
}

func (hastePotion *HastePotion) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting haste potion")
	hastePotion.CurrentCooldown = hastePotion.baseCooldown
	p.Am.TimerModifiers.HastePotion = hastePotion.buffDuration

	return &abilities.CastResult{
		Damage:           0,
		IsPhysical:       true,
		IsCriticalStrike: false,
		IsMiss:           true,
		CastTime:         0,
		DelayUntilCast:   0,
		OnGCD:            false,
	}
}

func (hastePotion HastePotion) Weight(p *player.Player) float64 {
	if hastePotion.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (hastePotion *HastePotion) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	hastePotion.CurrentCooldown = hastePotion.CurrentCooldown - timePassed
	return hastePotion.CurrentCooldown
}

func NewHastePotion() *HastePotion {
	return &HastePotion{
		baseCooldown:    120.0,
		buffDuration:    15,
		CurrentCooldown: 15,
	}
}
