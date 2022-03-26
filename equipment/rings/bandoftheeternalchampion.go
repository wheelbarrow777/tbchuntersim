package rings

import (
	"math"
	"tbchuntersim/abilities"
	"tbchuntersim/player"
	"tbchuntersim/util"

	log "github.com/sirupsen/logrus"
)

type BandOfTheEternalChampion struct {
	procPerMinute   float64
	buffDuration    float64
	procced         bool
	baseCooldown    float64
	CurrentCooldown float64
}

func (band *BandOfTheEternalChampion) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting BandOfTheEternalChampion")
	p.Am.TimerModifiers.BandOfTheEternalChampion = band.buffDuration
	band.CurrentCooldown = band.baseCooldown
	band.procced = false

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

func (dst BandOfTheEternalChampion) Weight(p *player.Player) float64 {
	if dst.procced {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (band *BandOfTheEternalChampion) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	// If it's off the cooldown, check if it procs
	if band.CurrentCooldown <= 0 {
		if opts.LastWasAHit {
			procChance := p.Equipment.Ranged.Speed * band.procPerMinute / 60
			if util.RollDice(procChance) {
				band.procced = true
			}
		}
	} else {
		// It's on the internal cooldown, remove the time that has passed
		timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
		band.CurrentCooldown = band.CurrentCooldown - timePassed
	}

	return band.CurrentCooldown
}

func NewBand() *BandOfTheEternalChampion {
	return &BandOfTheEternalChampion{
		procPerMinute: 1,
		buffDuration:  10,
		baseCooldown:  60,
	}
}
