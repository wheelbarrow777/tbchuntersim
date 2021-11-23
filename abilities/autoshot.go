package abilities

import (
	"math"
	"tbchuntersim/player"
	"tbchuntersim/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AutoShot struct {
	BaseCastTime    float64
	CurrentCooldown float64
}

var muteLog bool

// This number comes from the sixx sim
const asWeightModifier = 1.4

func (as *AutoShot) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	if opts.CastLast {
		as.CurrentCooldown = math.Max(0, p.RealSpeed()-as.castTime(p))
	} else {
		as.CurrentCooldown = math.Max(0, as.CurrentCooldown-(opts.ItTimeRev(0)-opts.ItTimeRev(2)))
	}

	return as.CurrentCooldown
}

func (as AutoShot) Cast(p *player.Player) *CastResult {
	if !muteLog {
		log.Debug("Casting auto shot")
	}

	ret := CastResult{
		CastTime:       as.castTime(p),
		OnGCD:          false,
		DelayUntilCast: as.CurrentCooldown,
		IsPhysical:     true,
		IsMiss:         false,
	}

	// Is it a miss?
	if util.RollDice(p.MissChance()) {
		ret.IsMiss = true
		return &ret
	}

	// We've hit!
	weapon := p.Equipment.Ranged
	unmodifiedDamage := weapon.AmmoDamage() + p.EffectiveAP()*weapon.Speed/14.0 + float64(weapon.Scope.Damage) + weapon.AverageDamage()

	// Is it a crit?
	if util.RollDice(p.CritChance()) {
		ret.IsCriticalStrike = true
		unmodifiedDamage = unmodifiedDamage * p.RangeCritDamageModifier()
	}

	ret.Damage = unmodifiedDamage * FOCUSED_FIRE_MOD
	return &ret
}

func (as AutoShot) Weight(p *player.Player) float64 {
	avgDmg := as.calcAvgDamage(p)
	a := avgDmg / (as.CurrentCooldown*p.RealSpeed()*asWeightModifier + as.castTime(p))
	return a
}

func (as AutoShot) castTime(p *player.Player) float64 {
	return as.BaseCastTime / p.TotalHaste()
}

func (as AutoShot) calcAvgDamage(p *player.Player) float64 {
	numIt := viper.GetInt("average-damage-iterations")
	if numIt == 0 {
		panic("numIt == 0")
	}
	dmgTotal := 0.0

	muteLog = true

	for i := 0; i < numIt; i++ {
		r := as.Cast(p)
		dmgTotal += r.Damage
	}

	muteLog = false
	return dmgTotal / float64(numIt)
}

func NewAutoShot() *AutoShot {
	as := AutoShot{
		BaseCastTime: 0.5,
	}

	return &as
}
