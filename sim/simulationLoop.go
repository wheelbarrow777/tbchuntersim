package sim

import (
	"math"
	"math/rand"
	"sort"
	"tbchuntersim/abilities"
	"tbchuntersim/consumables/potions"
	"tbchuntersim/equipment/trinkets"
	"tbchuntersim/player"
	"tbchuntersim/preset"
	"tbchuntersim/util"
	"time"

	log "github.com/sirupsen/logrus"
)

type simAbility struct {
	Ability     abilities.Ability
	Name        string
	Weight      float64
	NumHits     int
	TotalDamage float64
}

func newSimAbility(ability abilities.Ability, name string) simAbility {
	return simAbility{
		Ability:     ability,
		Weight:      0,
		TotalDamage: 0,
		NumHits:     0,
		Name:        name,
	}
}

func RunSimulationLoop(opts preset.SimOptions, p player.Player) *LoopResult {
	rand.Seed(time.Now().UnixMicro())

	var startTime float64 = 0
	var stopTime float64 = 0
	//var totalDamage float64 = 0
	iterationTimes := []float64{0, 0, 0, 0}
	gcdTracker := []float64{0, 0}

	var lastHadGCD bool = false // This variable is probably not needed
	var lastWasHit bool = false
	var lastWasCrit bool = false

	simRes := LoopResult{
		Ability: make(map[string]AbilityDetails),
	}

	// Initiate player
	abilityPriority := []simAbility{
		newSimAbility(abilities.NewAutoShot(), "AutoShot"),
		newSimAbility(abilities.NewSteadyShot(), "SteadyShot"),
		newSimAbility(abilities.NewTBW(), "The Beast Within"),
		newSimAbility(abilities.NewQuickShots(), "Quickshots"),
		newSimAbility(abilities.NewRapidFire(), "Rapid Fire"),
		newSimAbility(abilities.NewBloodlust(p.PlayerBuffs.Bloodlust), "Bloodlust"),
		newSimAbility(abilities.NewKillCommand(), "Kill Command"),
	}

	// Add abilities based on items
	if p.Equipment.TrinketOne.Name == "dragonspine trophy" || p.Equipment.TrinketTwo.Name == "dragonspine trophy" {
		abilityPriority = append(abilityPriority, newSimAbility(trinkets.NewDST(), "DST"))
	}
	if p.Equipment.TrinketOne.Name == "bloodlust brooch" || p.Equipment.TrinketTwo.Name == "bloodlust brooch" {
		abilityPriority = append(abilityPriority, newSimAbility(trinkets.NewBloodlustBrooch(), "Bloodlust Brooch"))
	}
	if p.ActivatedConsumables.LeatherworkingDrums {
		abilityPriority = append(abilityPriority, newSimAbility(potions.NewLeatherworkingDrums(), "Leatherworking Drums"))
	}
	if p.ActivatedConsumables.HastePotion {
		abilityPriority = append(abilityPriority, newSimAbility(potions.NewHastePotion(), "Haste Potion"))
	}
	if p.ActivatedConsumables.ManaPotion {
		abilityPriority = append(abilityPriority, newSimAbility(potions.NewSuperManaPotion(), "Super Mana Potion"))
	}

	// Adjust armor based on debuffs
	opts.TargetArmor = p.TargetDebuffs.EffectiveArmor(opts.TargetArmor)

	currentIteration := 0
	for stopTime < opts.SimDuration {

		// if p.CurrentMana < 200 {
		// 	util.NotifyLowMana()
		// }

		if currentIteration != 0 {

			// Calculate the cooldown of all abilities for this iteration
			for i, ability := range abilityPriority {
				ability.Ability.CalcCooldown(&p, &abilities.CalcCooldownOpts{
					CastLast:       i == 0,
					IterationTimes: iterationTimes,
					GCDTimes:       gcdTracker,
					LastHadGCD:     lastHadGCD,
					LastWasAHit:    lastWasHit,
					LastWasACrit:   lastWasCrit,
				})
			}

			// Calculate all the weights
			for i, a := range abilityPriority {
				abilityPriority[i].Weight = a.Ability.Weight(&p)
			}

			// Sort the abilities list by weight
			sort.Slice(abilityPriority, func(i, j int) bool {
				return abilityPriority[i].Weight > abilityPriority[j].Weight
			})
		}

		// Perform the first ability (The ability with the highest weight)
		castResult := abilityPriority[0].Ability.Cast(&p)

		// Calculate the real damage
		if castResult.IsPhysical {
			realDamageDealt := util.CalculateReducedArmorDamage(castResult.Damage, opts.TargetArmor)

			// Apply damage effects
			if p.TargetDebuffs.BloodFrenzy.Active {
				realDamageDealt = realDamageDealt * 1.04
			}

			// Apply ferocius inspiration
			for i := 0; i < p.PlayerBuffs.FerociousInspiration.Value+1; i++ {
				dmgModifier := 0.0
				if i == 0 {
					// The first check is the player itself. Only apply the benefit if the talent has been selected
					dmgModifier = 1.0 + 0.01*float64(p.Talents.BM.FerociousInspiration)
				} else {
					dmgModifier = 1.03
				}
				if util.RollDice(p.PlayerBuffs.FerociousInspiration.Uptime) {
					realDamageDealt = realDamageDealt * dmgModifier
				}
			}

			abilityPriority[0].NumHits++
			abilityPriority[0].TotalDamage += realDamageDealt
			simRes.Damage = append(simRes.Damage, realDamageDealt)
		} else {
			panic("non physical spells not implemented")
		}

		simRes.RangedAttackSpeed = append(simRes.RangedAttackSpeed, p.RealSpeed())

		// Calculate the times for the next iteration

		startTime = stopTime + castResult.DelayUntilCast
		stopTime = startTime + castResult.CastTime
		iterationTimes = append(iterationTimes, startTime, stopTime)
		if castResult.OnGCD {
			gcdTracker = append(gcdTracker, startTime+1.5)
		} else if startTime < gcdTracker[0] {
			gcdTracker = append(gcdTracker, gcdTracker[0])
		} else {
			gcdTracker = append(gcdTracker, 0)
		}
		lastHadGCD = castResult.OnGCD
		lastWasHit = !castResult.IsMiss
		lastWasCrit = castResult.IsCriticalStrike
		currentIteration++

		// Reduce the duration of all modifiers
		p.Am.ReduceModifierTime(stopTime - iterationTimes[len(iterationTimes)-3])

		// Apply MP5
		if stopTime > 5.0*math.Ceil(iterationTimes[len(iterationTimes)-3]/5.0) {
			p.ApplyMP5()
		}

		// Apply JOW
		if p.TargetDebuffs.JudgementOfWisdom.Active {
			// Only apply the bonus if the cast was a damaging ability
			if castResult.Damage > 0 {
				if util.RollDice(0.5) {
					p.AddMana(74)
				}
			}
		}

		// Log the time and current mana
		simRes.Time = append(simRes.Time, stopTime)
		simRes.Mana = append(simRes.Mana, p.CurrentMana)

		log.WithFields(log.Fields{
			"startTime": startTime,
			"stopTime":  stopTime,
			"speed":     p.Equipment.Ranged.Speed / p.TotalHaste(),
			"mana":      p.CurrentMana,
		}).Trace("Iteration Complete")

	}

	// Populate the return struct
	for _, a := range abilityPriority {
		simRes.Ability[a.Name] = AbilityDetails{
			TotalDamage: a.TotalDamage,
			NumHits:     a.NumHits,
		}
	}

	return &simRes
}
