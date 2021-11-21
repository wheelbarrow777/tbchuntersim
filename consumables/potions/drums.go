package potions

import (
	"math"
	"tbchuntersim/abilities"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type LeatherworkingDrums struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
}

func (drums *LeatherworkingDrums) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting drums")
	drums.CurrentCooldown = drums.baseCooldown
	p.Am.TimerModifiers.Drums = drums.buffDuration

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

func (drums LeatherworkingDrums) Weight(p *player.Player) float64 {
	if drums.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (drums *LeatherworkingDrums) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	drums.CurrentCooldown = drums.CurrentCooldown - timePassed
	return drums.CurrentCooldown
}

func NewLeatherworkingDrums() *LeatherworkingDrums {
	return &LeatherworkingDrums{
		baseCooldown:    120.0,
		buffDuration:    30.0,
		CurrentCooldown: 15.0,
	}
}
