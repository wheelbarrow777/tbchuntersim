package trinkets

import (
	"math"
	"tbchuntersim/abilities"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type BerserkersCall struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
}

func (bb *BerserkersCall) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting Berserkers Call")
	bb.CurrentCooldown = bb.baseCooldown
	p.Am.TimerModifiers.BerserkersCall = bb.buffDuration

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

func (bb BerserkersCall) Weight(p *player.Player) float64 {
	if bb.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (bb *BerserkersCall) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	bb.CurrentCooldown = bb.CurrentCooldown - timePassed
	return bb.CurrentCooldown
}

func NewBerserkersCall() *BerserkersCall {
	return &BerserkersCall{
		baseCooldown:    120.0,
		buffDuration:    20,
		CurrentCooldown: 15,
	}
}
