package player

import (
	"huntsim/consumables"
	"huntsim/equipment"
	"math"
	"strings"
)

type Race struct {
	Name string

	Agility   int
	Stamina   int
	Intellect int
	Spirit    int
	Strength  int

	Health    int
	Mana      int
	PetDamage float64
}

var (
	Orc = Race{
		Name:      "orc",
		Agility:   148,
		Stamina:   110,
		Intellect: 74,
		Spirit:    86,
		Strength:  67,
		Health:    3488,
		Mana:      3253,
		PetDamage: 1.05,
	}
	// TODO: Add troll crit modifier
	Troll = Race{
		Name:      "troll",
		Agility:   153,
		Stamina:   109,
		Intellect: 73,
		Spirit:    84,
		Strength:  65,
		Health:    3488,
		Mana:      3253,
		PetDamage: 1,
	}
)

type Player struct {
	Race Race

	Am ActiveModifiers

	PlayerBuffs          PlayerBuffs
	TargetDebuffs        TargetDebuffs
	StaticConsumeables   consumables.StaticConsumables
	ActivatedConsumables consumables.ActivatedConsumables

	MaxMana     float64
	CurrentMana float64

	Equipment equipment.Equipment
	Talents   Talents

	strength  int
	agility   int
	stamina   int
	intellect int
	spirit    int

	meleeAttackPower  int
	rangedAttackPower int
	critRating        int
	hitRating         int
	mp5               int
	armorPenetration  int
	hasteRating       float64

	critChanceMemoized struct {
		memoized bool
		value    float64
	}
}

type PlayerConfig struct {
	Race                 string
	Talents              Talents
	PlayerBuffs          PlayerBuffs
	TargetDebuffs        TargetDebuffs
	Equipment            equipment.Equipment
	StaticConsumables    consumables.StaticConsumables
	ActivatedConsumables consumables.ActivatedConsumables
}

func (p Player) RangeCritDamageModifier() float64 {
	// Mortal Shots
	mortalBonus := (1 + float64(p.Talents.MM.MortalShots)*0.06)

	// Earthstorm Diamond
	earthstormBonus := 0.0
	for _, gem := range p.Equipment.Helm.Gems.SlottedGems {
		if gem.Name == "relentless earthstorm diamond" {
			earthstormBonus = 1.03
			break
		}
	}

	return 1 + mortalBonus*(2*earthstormBonus-1)
}

func (p Player) CritChance() float64 {
	if !p.critChanceMemoized.memoized {
		agilityCrit := float64(p.agility) / AGILITY_TO_CRIT / 100
		critRatingCrit := float64(p.critRating) / CRIT_RATING_RATIO / 100
		lethalShotsCrit := float64(p.Talents.MM.LethalShots) * LETHAL_SHOTS_MODIFIER
		leaderOfThePackCrit := 0.0
		if p.PlayerBuffs.LeaderOfThePack.Active {
			leaderOfThePackCrit = 0.05
		}
		crit_supression := 0.048
		p.critChanceMemoized.value = agilityCrit + critRatingCrit + lethalShotsCrit + leaderOfThePackCrit - 0.0153 - crit_supression // TODO: Figure out these magic numbers
		if p.TargetDebuffs.JudgementOfTheCrusader.Active {
			p.critChanceMemoized.value += 0.03
		}
		p.critChanceMemoized.memoized = true
	}
	return p.critChanceMemoized.value
}

func (p Player) MissChance() float64 {
	base := 0.92

	if p.TargetDebuffs.FaeriFire.Active && p.TargetDebuffs.FaeriFire.Improved {
		base += IMP_FAERI_FIRE_HIT_VALUE
	}

	gear := float64(p.hitRating) / HIT_RATING_RATION / 100
	hitAdjustment := -0.01
	return 1 - math.Min(1, base+gear+hitAdjustment)
}

func (p Player) huntersMarkBonus() float64 {
	// TODO: Implement hunters mark bonus scaling
	return HUNTERS_MARK_AP_BONUS
}

func (p Player) EffectiveAP() float64 {
	base := float64(p.rangedAttackPower) + p.huntersMarkBonus()

	if p.Am.TimerModifiers.BloodlustBrooch > 0 {
		base += 278
	}

	if p.TargetDebuffs.ExposeWeakness.Active {
		base += p.TargetDebuffs.ExposeWeakness.Uptime * float64(p.TargetDebuffs.ExposeWeakness.Value)
	}

	return base
}

func (p *Player) calculateStats() {
	// Calculate the Stats
	// Base (Gear & Race)
	strength := float64(p.Equipment.Strength() + p.Race.Strength)
	agility := float64(p.Equipment.Agility() + p.Race.Agility)
	intellect := float64(p.Equipment.Intellect() + p.Race.Intellect)
	stamina := float64(p.Equipment.Stamina() + p.Race.Stamina)
	spirit := float64(p.Equipment.Spirit() + p.Race.Spirit)
	meleeAttackPower := float64(p.Equipment.AttackPower())
	rangedAttackPower := meleeAttackPower
	critRating := float64(p.Equipment.CritRating())
	hitRating := float64(p.Equipment.HitRating())
	mp5 := float64(p.Equipment.MP5())
	armorPenetration := float64(p.Equipment.ArmorPenetration())

	// Buffs
	buffs := p.PlayerBuffs
	// -- Gift of the Wild
	if buffs.GiftOfTheWild.Active {
		bonus := 1.0
		if p.PlayerBuffs.GiftOfTheWild.Improved {
			bonus = 1.35
		}
		strength += math.Floor(GIFT_OF_THE_WILD_STAT_VALUE * bonus)
		agility += math.Floor(GIFT_OF_THE_WILD_STAT_VALUE * bonus)
		intellect += math.Floor(GIFT_OF_THE_WILD_STAT_VALUE * bonus)
		stamina += math.Floor(GIFT_OF_THE_WILD_STAT_VALUE * bonus)
		spirit += math.Floor(GIFT_OF_THE_WILD_STAT_VALUE * bonus)
	}

	// -- BoM
	if buffs.BlessingOfMight.Active {
		bonus := 1.0
		if p.PlayerBuffs.BlessingOfMight.Improved {
			bonus = 1.2
		}
		meleeAttackPower += math.Floor(BOM_BONUS_AP * bonus)
		rangedAttackPower += math.Floor(BOM_BONUS_AP * bonus)
	}

	// -- BoW
	if buffs.BlessingOfWisdom.Active {
		bonus := 1.0
		if buffs.BlessingOfWisdom.Improved {
			bonus = 1.2
		}
		mp5 += math.Floor(BOW_MP5 * bonus)
	}

	// -- Battle Shout
	if buffs.BattleShout.Active {
		bonus := 1.0
		if buffs.BattleShout.Improved {
			bonus = 1.25
		}
		meleeAttackPower += math.Floor(BATTLE_SHOUT_MELEE_AP * bonus)
	}

	// -- Trueshot Aura
	if buffs.TrueShot.Active || p.Talents.MM.TrueshotAura > 0 {
		meleeAttackPower += TRUESHOT_AURA_AP
		rangedAttackPower += TRUESHOT_AURA_AP
	}

	// -- Leader of the Pack
	// This is applied at the crit calc level

	// -- Grace of Air Totem
	if buffs.GraceOfAirTotem.Active {
		bonus := 1.0
		if buffs.GraceOfAirTotem.Improved {
			bonus = 1.15
		}
		agility += math.Floor(GRACE_OF_AIR_VALUE * bonus)
	}

	// -- Strength of Earth
	if buffs.StrengthOfEarthTotem.Active {
		bonus := 1.0
		if buffs.StrengthOfEarthTotem.Improved {
			bonus = 1.15
		}
		strength += math.Floor(STRENGTH_OF_EARTH_VALUE * bonus)
	}

	// -- Mana Spring Totem
	if buffs.ManaSpringTotem.Active {
		mp5 += MANASPRING_MP5_VALUE
	}

	// -- Windfury Totem
	// TODO: Implement with melee weaving patch
	if buffs.WindfuryTotem.Active {
		panic("windfury totem is currently unsupported")
	}

	// -- Arcane Brilliance
	if buffs.ArcaneBrilliance.Active {
		intellect += ARCANE_BRILLIANCE_VALUE
	}

	// -- Prayer of Fortitude
	if buffs.PrayerOfFortitude.Active {
		bonus := 1.0
		if buffs.PrayerOfFortitude.Improved {
			bonus = 1.3
		}
		stamina += math.Floor(PRAYER_OF_FORTITUDE_VALUE * bonus)
	}

	// -- Blood Pact
	if buffs.BloodPact.Active {
		bonus := 1.0
		if buffs.BloodPact.Improved {
			bonus = 1.3
		}
		stamina += BLOOD_PACT_VALUE * bonus
	}

	// -- Braided eternium chain
	if buffs.BraidedEterniumChain.Active {
		critRating += 28
	}

	// Consumables
	sumCons := p.StaticConsumeables.SummedStatsPlayer()
	agility += float64(sumCons.Agility)
	strength += float64(sumCons.Strength)
	intellect += float64(sumCons.Intellect)
	critRating += float64(sumCons.CritRating)
	hitRating += float64(sumCons.HitRating)
	spirit += float64(sumCons.Spirit)
	meleeAttackPower += float64(sumCons.AttackPower)
	rangedAttackPower += float64(sumCons.AttackPower)
	mp5 += float64(sumCons.MP5)
	stamina += float64(sumCons.Stamina)

	// -- Apply oil
	// Todo: add support for other oils
	critRating += float64(p.Equipment.MainHand.Oil.CritRating)
	critRating += float64(p.Equipment.OffHand.Oil.CritRating)

	// Blessing of Kings
	if buffs.BlessingOfKings.Active {
		agility = agility * BOK_MODIFIER
		strength = strength * BOK_MODIFIER
		intellect = intellect * BOK_MODIFIER
		stamina = stamina * BOK_MODIFIER
		spirit = spirit * BOK_MODIFIER
	}

	// Talents
	// -- This must be applied at the end
	rangedAttackPower += float64(p.Talents.MM.CarefulAim) * 0.15 * intellect

	// Apply all the finished calculated stats to the player
	p.agility = int(math.Floor(agility))
	p.strength = int(math.Floor(strength))
	p.intellect = int(math.Floor(intellect))
	p.stamina = int(math.Floor(stamina))
	p.spirit = int(math.Floor(spirit))

	// No clue about these magic numbers. But they are needed for the final number to be correct
	p.meleeAttackPower = int(math.Floor(meleeAttackPower)) + p.agility + p.strength + 120
	p.rangedAttackPower = int(math.Floor(rangedAttackPower)) + p.agility - 10 + 140 + 155

	p.critRating = int(math.Floor(critRating))
	p.hitRating = int(math.Floor(hitRating))
	p.mp5 = int(math.Floor(mp5))
	p.armorPenetration = int(math.Floor(armorPenetration))

	// Calculate max mana
	p.MaxMana = float64(p.Race.Mana) + (float64(p.intellect)-10.0)*15
	p.CurrentMana = p.MaxMana
}

func (p *Player) ApplyMP5() {
	p.AddMana(float64(p.mp5))
}

func (p *Player) AddMana(value float64) {
	p.CurrentMana = math.Min(p.MaxMana, p.CurrentMana+value)

}

func NewPlayer(config *PlayerConfig) *Player {
	p := Player{
		PlayerBuffs:          config.PlayerBuffs,
		TargetDebuffs:        config.TargetDebuffs,
		Equipment:            config.Equipment,
		Talents:              config.Talents,
		StaticConsumeables:   config.StaticConsumables,
		ActivatedConsumables: config.ActivatedConsumables,
	}

	if strings.ToLower(config.Race) == "orc" {
		p.Race = Orc
	} else if strings.ToLower(config.Race) == "troll" {
		p.Race = Troll
	} else {
		panic("invalid race")
	}

	p.calculateStats()

	return &p
}
