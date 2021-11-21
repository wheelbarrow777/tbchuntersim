package abilities

import (
	"math"
	"tbchuntersim/player"
	"tbchuntersim/util"

	log "github.com/sirupsen/logrus"
)

type QuickShots struct {
	procChance      float64
	buffDuration    float64
	procced         bool
	CurrentCooldown float64
}

func (qs *QuickShots) Cast(p *player.Player) *CastResult {
	log.Debug("Casting Quickshots")
	p.Am.TimerModifiers.QuickShots = qs.buffDuration
	qs.CurrentCooldown = qs.buffDuration
	qs.procced = false

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

func (qs QuickShots) Weight(p *player.Player) float64 {
	if qs.procced {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (qs *QuickShots) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	// If it's off the cooldown, check if it procs
	if qs.CurrentCooldown <= 0 {
		if opts.LastWasAHit {
			if util.RollDice(qs.procChance) {
				qs.procced = true
			}
		}
	} else {
		// It's on the internal cooldown, remove the time that has passed
		timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
		qs.CurrentCooldown = qs.CurrentCooldown - timePassed
	}

	return qs.CurrentCooldown
}

func NewQuickShots() *QuickShots {
	return &QuickShots{
		procChance:   0.1,
		buffDuration: 12,
	}
}
