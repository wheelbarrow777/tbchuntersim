package potions

import (
	"math"
	"math/rand"
	"tbchuntersim/abilities"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type ManaPotion struct {
	baseCooldown    float64
	manaGainMin     int
	manaGainmax     int
	CurrentCooldown float64
}

func (manaPotion ManaPotion) calculateManaGain() int {
	return manaPotion.manaGainMin + rand.Intn(manaPotion.manaGainmax-manaPotion.manaGainMin)
}

func (manaPotion *ManaPotion) Cast(p *player.Player) *abilities.CastResult {
	log.Debug("Casting haste potion")
	manaPotion.CurrentCooldown = manaPotion.baseCooldown

	p.CurrentMana = math.Min(p.MaxMana, p.CurrentMana+float64(manaPotion.calculateManaGain()))

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

func (manaPotion ManaPotion) Weight(p *player.Player) float64 {

	// If pot is in cooldown, can't cast
	if manaPotion.CurrentCooldown > 0 {
		return math.Inf(-1)
	}

	// If we can get all the mana form the pot, use it
	if int(p.CurrentMana)+manaPotion.manaGainmax < int(p.MaxMana) {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (manaPotion *ManaPotion) CalcCooldown(p *player.Player, opts *abilities.CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	manaPotion.CurrentCooldown = manaPotion.CurrentCooldown - timePassed
	return manaPotion.CurrentCooldown
}

func NewSuperManaPotion() *ManaPotion {
	return &ManaPotion{
		baseCooldown:    120.0,
		manaGainMin:     1800,
		manaGainmax:     3000,
		CurrentCooldown: 15,
	}
}
