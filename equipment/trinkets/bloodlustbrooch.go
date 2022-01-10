package trinkets

import (
	"math"
	"tbchuntersim/abilities"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type BloodlustBrooch struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
}

func (bb *BloodlustBrooch) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting Bloodlust Brooch")
	bb.CurrentCooldown = bb.baseCooldown
	p.Am.TimerModifiers.BloodlustBrooch = bb.buffDuration

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

func (bb BloodlustBrooch) Weight(p *player.Player) float64 {
	if bb.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (bb *BloodlustBrooch) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	bb.CurrentCooldown = bb.CurrentCooldown - timePassed
	return bb.CurrentCooldown
}

func NewBloodlustBrooch() *BloodlustBrooch {
	return &BloodlustBrooch{
		baseCooldown:    120.0,
		buffDuration:    20,
		CurrentCooldown: 15,
	}
}
