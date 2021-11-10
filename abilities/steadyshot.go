package abilities

import (
	"huntsim/player"
	"huntsim/util"
	"math"

	log "github.com/sirupsen/logrus"
)

type SteadyShot struct {
	BaseCastTime    float64
	CurrentCooldown float64
	ManaCost        float64
}

func (ss SteadyShot) Cast(p *player.Player) *CastResult {
	log.Debug("Casting steady shot")
	ret := ss.calcDamage(p)

	// Take away the mana
	if p.CurrentMana < ss.ManaCost {
		ret.Damage = 0
	} else {
		p.CurrentMana -= ss.ManaCost
	}

	return ret

}

func (ss SteadyShot) calcDamage(p *player.Player) *CastResult {
	ret := CastResult{
		Damage:         0,
		CastTime:       ss.castTime(p),
		OnGCD:          true,
		DelayUntilCast: ss.CurrentCooldown,
		IsPhysical:     true,
		IsMiss:         false,
	}

	// Is it a miss?
	if util.RollDice(p.MissChance()) {
		return &ret
	}

	// We've hit!
	unmodifiedDamage := (p.Equipment.Ranged.AverageDamage()*2.8/p.Equipment.Ranged.Speed + p.EffectiveAP()*0.2 + 150.0)

	// Is it a crit?
	if util.RollDice(p.CritChance()) {
		ret.IsCriticalStrike = true
		unmodifiedDamage = unmodifiedDamage * p.RangeCritDamageModifier()
	}

	ret.Damage = unmodifiedDamage * FOCUSED_FIRE_MOD

	if p.Am.TimerModifiers.TBW > 0 {
		ret.Damage = ret.Damage * TBW_MOD
	}

	return &ret
}

func (ss *SteadyShot) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	if opts.CastLast {
		ss.CurrentCooldown = math.Max(opts.GCDTmeRev(0)-opts.ItTimeRev(0), 0)
	} else {
		gcdOffset := 0.0
		if opts.LastHadGCD {
			gcdOffset = opts.GCDTmeRev(0) - opts.ItTimeRev(0)
		}
		ss.CurrentCooldown = math.Max(ss.CurrentCooldown-opts.ItTimeRev(0)-opts.ItTimeRev(2)+gcdOffset, 0)
	}
	return ss.CurrentCooldown
}

func (ss SteadyShot) Weight(p *player.Player) float64 {
	if p.CurrentMana < ss.ManaCost {
		return math.Inf(-1)
	}
	avgDmg := ss.calcAvgDamage(p)

	a := avgDmg / (math.Max(ss.CurrentCooldown, 0)*p.RealSpeed()*1.2 + ss.castTime(p))
	return a
}

func (ss SteadyShot) calcAvgDamage(p *player.Player) float64 {
	numIt := 500
	dmgTotal := 0.0
	for i := 0; i < numIt; i++ {
		r := ss.calcDamage(p)
		dmgTotal += r.Damage
	}
	return dmgTotal / float64(numIt)
}

func (ss SteadyShot) castTime(p *player.Player) float64 {
	return ss.BaseCastTime / p.TotalHaste()
}

func NewSteadyShot() *SteadyShot {
	return &SteadyShot{
		BaseCastTime: 1.5,
		ManaCost:     99.0,
	}
}
