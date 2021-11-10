package trinkets

import (
	"huntsim/abilities"
	"huntsim/player"
	"huntsim/util"
	"math"

	log "github.com/sirupsen/logrus"
)

type DragonspineTrophy struct {
	procPerMinute   float64
	buffDuration    float64
	procced         bool
	baseCooldown    float64
	CurrentCooldown float64
}

func (dst *DragonspineTrophy) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting DST")
	p.Am.TimerModifiers.DST = dst.buffDuration
	dst.CurrentCooldown = dst.baseCooldown
	dst.procced = false

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

func (dst DragonspineTrophy) Weight(p *player.Player) float64 {
	if dst.procced {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (dst *DragonspineTrophy) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	// If it's off the cooldown, check if it procs
	if dst.CurrentCooldown <= 0 {
		if opts.LastWasAHit {
			procChance := p.Equipment.Ranged.Speed * dst.procPerMinute / 60
			if util.RollDice(procChance) {
				dst.procced = true
			}
		}
	} else {
		// It's on the internal cooldown, remove the time that has passed
		timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
		dst.CurrentCooldown = dst.CurrentCooldown - timePassed
	}

	return dst.CurrentCooldown
}

func NewDST() *DragonspineTrophy {
	return &DragonspineTrophy{
		procPerMinute: 1,
		buffDuration:  10,
		baseCooldown:  20,
	}
}
