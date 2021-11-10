package abilities

import (
	"huntsim/player"
	"math"

	log "github.com/sirupsen/logrus"
)

type RapidFire struct {
	baseCooldown    float64
	buffDuration    float64
	CurrentCooldown float64
	ManaCost        float64
}

func (rf *RapidFire) Cast(p *player.Player) *CastResult {
	log.Debug("Casting Rapid Fire")
	rf.CurrentCooldown = rf.baseCooldown
	p.Am.TimerModifiers.RapidFire = rf.buffDuration

	// Take away the mana
	if p.CurrentMana < rf.ManaCost {
		panic("tried to cast rapid fire without mana")
	} else {
		p.CurrentMana -= rf.ManaCost
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

func (rf RapidFire) Weight(p *player.Player) float64 {
	if p.CurrentMana < rf.ManaCost {
		return math.Inf(-1)
	}
	if rf.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (rf *RapidFire) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	rf.CurrentCooldown = rf.CurrentCooldown - timePassed
	return rf.CurrentCooldown
}

func NewRapidFire() *RapidFire {
	return &RapidFire{
		baseCooldown:    180.0,
		buffDuration:    15,
		CurrentCooldown: 15,
	}
}
