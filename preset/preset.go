package preset

import (
	"tbchuntersim/consumables"
	"tbchuntersim/player"
)

type SimulationPreset struct {
	Race                 string
	Equipment            eq
	Consumables          consums
	ActivatedConsumables consumables.ActivatedConsumables
	Buffs                buffs
	TargetDebuffs        player.TargetDebuffs
	Talents              player.Talents
	Options              SimOptions
}

type itemBase struct {
	Name    string
	Gems    []string
	Enchant string
}

type weapon struct {
	itemBase
	Oil string
}

type rangedWeapon struct {
	itemBase
	Scope string
}

type eq struct {
	Helm       itemBase
	Neck       itemBase
	Shoulders  itemBase
	Cloak      itemBase
	Chest      itemBase
	Bracers    itemBase
	Gloves     itemBase
	Belt       itemBase
	Pants      itemBase
	Boots      itemBase
	RingOne    itemBase
	RingTwo    itemBase
	TrinketOne itemBase
	TrinketTwo itemBase
	MainHand   weapon
	OffHand    weapon
	Ranged     rangedWeapon
	Quiver     string
	AmmoDPS    float64
}

type consums struct {
	Food           string
	BattleElixir   string
	GuardianElixir string
	AgilityScroll  string
	StrengthScroll string

	PetFood           string
	PetScrollAgility  bool
	PetScrollStrenght bool
}
type buffs struct {
	BlessingOfKings      player.Buff
	BlessingOfMight      player.BuffWithImproved
	BlessingOfWisdom     player.BuffWithImproved
	BattleShout          player.BuffWithImproved
	TrueShotAura         bool
	LeaderOfThePack      player.BuffWithImproved
	GraceOfAirTotem      player.BuffWithImproved
	StrengthOfEarthTotem player.BuffWithImproved
	ManaSpringTotem      player.Buff
	// WindfuryTotem buff // Not supported
	ArcaneBrilliance          bool
	GiftOfTheWild             player.BuffWithImproved
	BloodLustCount            int
	FerociousInspirationCount struct {
		ExtraHunters int
		Uptime       float64
	}
	PrayerOfFortitude    player.BuffWithImproved
	BloodPact            player.BuffWithImproved
	BraidedEterniumChain bool
}

type SimOptions struct {
	SimDuration float64
	Latency     float64
	TargetArmor float64
}
