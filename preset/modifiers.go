package preset

import "tbchuntersim/player"

var (
	AllBuffsMod = SimulationPreset{
		Buffs: buffs{
			BlessingOfKings: player.Buff{
				Active: true,
			},
			BlessingOfMight: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			BlessingOfWisdom: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			BattleShout: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			TrueShotAura: true,
			LeaderOfThePack: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			GraceOfAirTotem: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			StrengthOfEarthTotem: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			ManaSpringTotem: player.Buff{
				Active: true,
			},
			ArcaneBrilliance: true,
			GiftOfTheWild: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			BloodLustCount: 1,
			FerociousInspirationCount: struct {
				ExtraHunters int
				Uptime       float64
			}{
				ExtraHunters: 2,
				Uptime:       0.94,
			},
			PrayerOfFortitude: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			BloodPact: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			BraidedEterniumChain: true,
		},
	}
	NoBuffsMod = SimulationPreset{
		Buffs: buffs{
			BlessingOfKings: player.Buff{
				Active: false,
			},
			BlessingOfMight: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			BlessingOfWisdom: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			BattleShout: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			TrueShotAura: false,
			LeaderOfThePack: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			GraceOfAirTotem: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			StrengthOfEarthTotem: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			ManaSpringTotem: player.Buff{
				Active: false,
			},
			ArcaneBrilliance: false,
			GiftOfTheWild: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			BloodLustCount: 0,
			FerociousInspirationCount: struct {
				ExtraHunters int
				Uptime       float64
			}{
				ExtraHunters: 0,
				Uptime:       0.94,
			},
			PrayerOfFortitude: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			BloodPact: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			BraidedEterniumChain: false,
		},
	}
	AllDebuffsMod = SimulationPreset{
		TargetDebuffs: player.TargetDebuffs{
			ImprovedHuntersMark: player.Buff{
				Active: true,
			},
			JudgementOfWisdom: player.Buff{
				Active: true,
			},
			JudgementOfTheCrusader: player.Buff{
				Active: true,
			},
			CurseOfElements: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			ExposeWeakness: player.BuffWithUptimeAndValue{
				BuffWithUptime: player.BuffWithUptime{
					Uptime: 0.92,
					Buff: player.Buff{
						Active: true,
					},
				},
				Value: 284,
			},
			SunderArmor: player.Buff{
				Active: true,
			},
			ExposeArmor: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			CurseOfRecklessness: player.Buff{
				Active: true,
			},
			FaeriFire: player.BuffWithImproved{
				Buff: player.Buff{
					Active: true,
				},
				Improved: true,
			},
			Misery: player.Buff{
				Active: true,
			},
			BloodFrenzy: player.Buff{
				Active: true,
			},
		},
	}
	NoDebuffsMod = SimulationPreset{
		TargetDebuffs: player.TargetDebuffs{
			ImprovedHuntersMark: player.Buff{
				Active: false,
			},
			JudgementOfWisdom: player.Buff{
				Active: false,
			},
			JudgementOfTheCrusader: player.Buff{
				Active: false,
			},
			CurseOfElements: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			ExposeWeakness: player.BuffWithUptimeAndValue{
				BuffWithUptime: player.BuffWithUptime{
					Uptime: 0.92,
					Buff: player.Buff{
						Active: false,
					},
				},
				Value: 284,
			},
			SunderArmor: player.Buff{
				Active: false,
			},
			ExposeArmor: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			CurseOfRecklessness: player.Buff{
				Active: false,
			},
			FaeriFire: player.BuffWithImproved{
				Buff: player.Buff{
					Active: false,
				},
				Improved: false,
			},
			Misery: player.Buff{
				Active: false,
			},
			BloodFrenzy: player.Buff{
				Active: false,
			},
		},
	}
	AllConsumablesMod = SimulationPreset{
		Consumables: consums{
			Food:              "Warp Burger",
			BattleElixir:      "Elixir of Major Agility",
			GuardianElixir:    "Elixir of Major Mageblood",
			AgilityScroll:     "Scroll of Agility V",
			StrengthScroll:    "Scroll of Strength V",
			PetFood:           "Kibler's Bits",
			PetScrollAgility:  true,
			PetScrollStrenght: true,
		},
	}
	NoConsumablesMod = SimulationPreset{
		Consumables: consums{
			Food:              "",
			BattleElixir:      "",
			GuardianElixir:    "",
			AgilityScroll:     "",
			StrengthScroll:    "",
			PetFood:           "",
			PetScrollAgility:  false,
			PetScrollStrenght: false,
		},
	}
)
