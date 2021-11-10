package abilities

import (
	"huntsim/player"
	"math"

	log "github.com/sirupsen/logrus"
)

type Bloodlust struct {
	numOfLusts      int
	buffDuration    float64
	CurrentCooldown float64
}

func (bl *Bloodlust) Cast(p *player.Player) *CastResult {
	log.Debug("Casting Bloodlust")

	r := CastResult{
		Damage:           0,
		IsPhysical:       true,
		IsCriticalStrike: false,
		IsMiss:           true,
		CastTime:         0,
		DelayUntilCast:   0,
		OnGCD:            false,
	}

	if bl.numOfLusts <= 0 {
		log.Warn("Tried to cast lust without any more charges")
		bl.CurrentCooldown = math.Inf(1)
		return &r
	} else {
		bl.CurrentCooldown = bl.buffDuration
		bl.numOfLusts--
	}

	p.Am.TimerModifiers.Bloodlust = bl.buffDuration

	return &r
}

func (bl Bloodlust) Weight(p *player.Player) float64 {
	if bl.numOfLusts <= 0 {
		return math.Inf(-1)
	}
	if bl.CurrentCooldown <= 0 {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (bl *Bloodlust) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	bl.CurrentCooldown = bl.CurrentCooldown - timePassed
	return bl.CurrentCooldown
}

func NewBloodlust(numLusts int) *Bloodlust {
	return &Bloodlust{
		numOfLusts:      numLusts,
		buffDuration:    40,
		CurrentCooldown: 15,
	}
}
