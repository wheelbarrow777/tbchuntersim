package player

const (
	STRENGTH_OF_EARTH_VALUE      = 86
	GRACE_OF_AIR_VALUE           = 77
	MANASPRING_MP5_VALUE         = 50
	GIFT_OF_THE_WILD_STAT_VALUE  = 14
	GIFT_OF_THE_WILD_ARMOR_VALUE = 340
	GIFT_OF_THE_WILD_RES_VALUE   = 25
	BOK_MODIFIER                 = 1.1
	BOM_BONUS_AP                 = 220
	BOW_MP5                      = 41
	BATTLE_SHOUT_MELEE_AP        = 305
	TRUESHOT_AURA_AP             = 125
	ARCANE_BRILLIANCE_VALUE      = 40
	PRAYER_OF_FORTITUDE_VALUE    = 79
	BLOOD_PACT_VALUE             = 66

	SUNDER_ARMOR_ARMOR_VALUE          = 520
	EXPOSED_ARMOR_BASE_VALUE          = 2050.0
	CURSE_OF_RECKLESSNESS_ARMOR_VALUE = 800
	FAERI_FIRE_ARMOR_VALUE            = 610
	IMP_FAERI_FIRE_HIT_VALUE          = 0.03
)

type Buff struct {
	Active bool
}

type BuffWithImproved struct {
	Buff
	Improved bool
}

type BuffWithUptime struct {
	Buff
	Uptime float64
}

type BuffWithUptimeAndValue struct {
	BuffWithUptime
	Value int
}

type TargetDebuffs struct {
	ImprovedHuntersMark    Buff
	JudgementOfWisdom      Buff
	JudgementOfTheCrusader Buff

	CurseOfElements     BuffWithImproved
	ExposeWeakness      BuffWithUptimeAndValue
	SunderArmor         Buff
	ExposeArmor         BuffWithImproved
	CurseOfRecklessness Buff
	FaeriFire           BuffWithImproved
	Misery              Buff
	BloodFrenzy         Buff
}

type PlayerBuffs struct {
	// Paladin
	BlessingOfKings  Buff
	BlessingOfMight  BuffWithImproved
	BlessingOfWisdom BuffWithImproved

	// Warrior
	BattleShout BuffWithImproved

	// Hunter
	TrueShot             Buff
	FerociousInspiration BuffWithUptimeAndValue

	// Shaman
	GraceOfAirTotem      BuffWithImproved
	StrengthOfEarthTotem BuffWithImproved
	ManaSpringTotem      Buff
	WindfuryTotem        Buff

	// Mage
	ArcaneBrilliance Buff

	// Druid
	GiftOfTheWild   BuffWithImproved
	LeaderOfThePack BuffWithImproved

	// Shaman
	Bloodlust int

	// Priest
	PrayerOfFortitude BuffWithImproved

	// Warlock
	BloodPact BuffWithImproved

	// Jewelcrafting
	BraidedEterniumChain Buff
}

func (t TargetDebuffs) EffectiveArmor(startingArmor float64) float64 {
	armor := startingArmor
	if t.SunderArmor.Active {
		armor -= 5 * SUNDER_ARMOR_ARMOR_VALUE
	}
	if t.ExposeArmor.Active {
		exposed := EXPOSED_ARMOR_BASE_VALUE
		if t.ExposeArmor.Improved {
			exposed = exposed * 1.5
		}
		armor -= exposed
	}
	if t.CurseOfRecklessness.Active {
		armor -= CURSE_OF_RECKLESSNESS_ARMOR_VALUE
	}
	if t.FaeriFire.Active {
		armor -= FAERI_FIRE_ARMOR_VALUE
	}

	return armor

}
