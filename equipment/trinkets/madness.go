package trinkets

import (
	"math"
	"tbchuntersim/abilities"
	"tbchuntersim/player"
	"tbchuntersim/util"
)

type MadnessOfTheBetrayer struct {
	procPerMinute float64
	buffDuration  float64
	procced       bool
}

func (madness *MadnessOfTheBetrayer) Cast(p *player.Player) *abilities.CastResult {
	p.Am.TimerModifiers.Madness = madness.buffDuration
	madness.procced = false

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

func (madness MadnessOfTheBetrayer) Weight(p *player.Player) float64 {
	if madness.procced {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (madness *MadnessOfTheBetrayer) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	// No internal cooldown
	if opts.LastWasAHit {
		procChance := p.Equipment.Ranged.Speed * madness.procPerMinute / 60
		if util.RollDice(procChance) {
			madness.procced = true
		}
	}

	return 0
}

func NewMadness() *MadnessOfTheBetrayer {
	return &MadnessOfTheBetrayer{
		procPerMinute: 1,
		buffDuration:  10,
	}
}
