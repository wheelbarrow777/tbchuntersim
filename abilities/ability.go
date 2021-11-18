package abilities

import "huntsim/player"

type CalcCooldownOpts struct {
	CastLast     bool
	LastHadGCD   bool
	LastWasAHit  bool
	LastWasACrit bool

	IterationTimes []float64
	GCDTimes       []float64
}

type CastResult struct {
	Damage float64

	IsPhysical       bool
	IsCriticalStrike bool
	IsMiss           bool

	CastTime       float64
	DelayUntilCast float64
	OnGCD          bool
}

type Ability interface {
	Cast(p *player.Player) *CastResult
	Weight(p *player.Player) float64
	CalcCooldown(p *player.Player, opts *CalcCooldownOpts) float64
}

func (opts CalcCooldownOpts) ItTimeRev(i int) float64 {
	return opts.IterationTimes[len(opts.IterationTimes)-1-i]
}

func (opts CalcCooldownOpts) GCDTmeRev(i int) float64 {
	return opts.GCDTimes[len(opts.GCDTimes)-1-i]
}
