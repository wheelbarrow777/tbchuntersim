package abilities

import (
	"math"
	"tbchuntersim/player"

	log "github.com/sirupsen/logrus"
)

type KillCommand struct {
	ManaCost        float64
	CurrentCooldown float64
	baseCooldown    float64
	procced         bool
}

func (kc *KillCommand) Cast(p *player.Player) *CastResult {
	log.Debug("Casting Kill Command")
	kc.CurrentCooldown = kc.baseCooldown

	// Take away the mana
	if p.CurrentMana < kc.ManaCost {
		panic("tried to cast kill command without mana")
	} else {
		p.CurrentMana -= kc.ManaCost
	}

	if p.Equipment.HasBeastLordFourSet() {
		p.Am.TimerModifiers.BeastLordArmorPen = 15.0
	}

	kc.procced = false

	return &CastResult{
		Damage:           0, // The damage is done by the pet and not the hunter
		IsPhysical:       true,
		IsCriticalStrike: false,
		IsMiss:           true,
		CastTime:         0,
		DelayUntilCast:   0,
		OnGCD:            false,
	}

}

func (kc KillCommand) Weight(p *player.Player) float64 {
	if p.CurrentMana < kc.ManaCost {
		return math.Inf(-1)
	}

	if kc.CurrentCooldown > 0 {
		return math.Inf(-1)
	}

	if kc.procced {
		return math.Inf(1)
	} else {
		return math.Inf(-1)
	}
}

func (kc *KillCommand) CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64 {
	timePassed := opts.ItTimeRev(0) - opts.ItTimeRev(2)
	kc.CurrentCooldown = kc.CurrentCooldown - timePassed
	if opts.LastWasACrit {
		kc.procced = true
	}

	return 0
}

func NewKillCommand() *KillCommand {
	return &KillCommand{
		ManaCost:        75,
		procced:         false,
		CurrentCooldown: 0,
		baseCooldown:    5,
	}
}
