package sim

import (
	"huntsim/abilities"
	"huntsim/config"
	"huntsim/equipment/trinkets"
	"huntsim/player"
	"huntsim/util"
	"math/rand"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

type simAbility struct {
	Ability     abilities.Ability
	Name        string
	Weight      float64
	NumHits     float64
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

func RunSimulationLoop(opts *config.SimOptions, p player.Player) float64 {
	var startTime float64 = 0
	var stopTime float64 = 0
	var totalDamage float64 = 0
	iterationTimes := []float64{0, 0, 0, 0}
	gcdTracker := []float64{0, 0}
	var lastHadGCD bool = false // This variable is probably not needed
	var lastWasHit bool = false

	// Initiate player
	abilityPriority := []simAbility{
		newSimAbility(abilities.NewAutoShot(), "AutoShot"),
		newSimAbility(abilities.NewSteadyShot(), "SteadyShot"),
		newSimAbility(abilities.NewTBW(), "The Beast Within"),
		newSimAbility(trinkets.NewDST(), "DST"), // TODO: Add dynamically based on which trinket is equipped
		newSimAbility(trinkets.NewBloodlustBrooch(), "Bloodlust Brooch"),
		newSimAbility(abilities.NewQuickShots(), "Quickshots"),
		newSimAbility(abilities.NewRapidFire(), "Rapid Fire"),
		//newSimAbility(abilities.NewBloodlust(1), "Bloodlust"),
	}

	rand.Seed(time.Now().UnixMicro())

	currentIteration := 0
	for stopTime < opts.SimDuration {
		if currentIteration != 0 {
			// TODO CHEAT
			p.CurrentMana = p.MaxMana

			// Calculate the cooldown of all abilities for this iteration
			for i, ability := range abilityPriority {
				ability.Ability.CalcCooldown(&p, &abilities.CalcCooldownOpts{
					CastLast:       i == 0,
					IterationTimes: iterationTimes,
					GCDTimes:       gcdTracker,
					LastHadGCD:     lastHadGCD,
					LastWasAHit:    lastWasHit,
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

		if castResult.IsPhysical {
			realDamageDealt := util.CalculateReducedArmorDamage(castResult.Damage, opts.TargetArmor)
			totalDamage += realDamageDealt
			abilityPriority[0].NumHits++
			abilityPriority[0].TotalDamage += realDamageDealt
		} else {
			panic("non physical spells not implemented")
		}

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
		currentIteration++

		// Reduce the duration of all modifiers
		//	log.Debugf("Reduced time by %f", stopTime-iterationTimes[len(iterationTimes)-3])
		p.Am.ReduceModifierTime(stopTime - iterationTimes[len(iterationTimes)-3])

		// Did anything proc?
		// - Kill Command
		if castResult.IsCriticalStrike {
			// Activate kill command
			if p.Equipment.HasBeastLordFourSet() {
				p.Am.BeastLordArmorPen = 15.0
			}
		}

		log.WithFields(log.Fields{
			"startTime": startTime,
			"stopTime":  stopTime,
			"speed":     p.Equipment.Ranged.Speed / p.TotalHaste(),
			"mana":      p.CurrentMana,
		}).Debug("Iteration Complete")

	}

	// for _, a := range abilityPriority {
	// 	if a.Name == "AutoShot" || a.Name == "SteadyShot" {
	// 		log.WithFields(log.Fields{
	// 			"Name":         a.Name,
	// 			"Total Damage": a.TotalDamage,
	// 			"Num of Hits":  a.NumHits,
	// 			"Average Hit":  a.TotalDamage / a.NumHits,
	// 			"Mana":         p.CurrentMana,
	// 		}).Info("Damage total")
	// 	}
	// }
	return totalDamage / stopTime
}
