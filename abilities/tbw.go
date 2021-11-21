package abilities

import (
	"math"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type TBW struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
	ManaCost        float64
}

func (tbw *TBW) Cast(p *player.Player) *CastResult {
	log.Debug("Casting TBW")
	tbw.CurrentCooldown = tbw.baseCooldown
	p.Am.TimerModifiers.TBW = tbw.buffDuration

	// Take away the mana
	if p.CurrentMana < tbw.ManaCost {
		panic("tried to cast tbw without mana")
	} else {
		p.CurrentMana -= tbw.ManaCost
	}

	return &CastResult{
		Damage:           0,
		IsPhysical:       true,
		IsCriticalStrike: false,
		IsMiss:           true,
		CastTime:         0,
		DelayUntilCast:   0,
		OnGCD:            false,
	}
}

func (tbw TBW) Weight(p *player.Player) float64 {
	if p.CurrentMana < tbw.ManaCost {
		return math.Inf(-1)
	}
	if tbw.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func NewTBW() *TBW {
	return &TBW{
		baseCooldown:    120.0,
		buffDuration:    18,
		CurrentCooldown: 15,
	}
}

func (tbw *TBW) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	tbw.CurrentCooldown = tbw.CurrentCooldown - timePassed
	return tbw.CurrentCooldown
}
