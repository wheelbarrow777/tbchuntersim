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
)

type Buff struct {
	Active   bool
	Improved bool
}

type TargetDebuffs struct {
	ImprovedHuntersMark    Buff
	JudgementOfWisdom      Buff
	JudgementOfTheCrusader Buff

	CurseOfElements     Buff
	SunderArmor         Buff
	ImprovedExposeArmor Buff
	CurseOfRecklessness Buff
	FaeriFire           Buff
	Misery              Buff
	BloodFrenzy         Buff
}

type PlayerBuffs struct {
	// Paladin
	BlessingOfKings  Buff
	BlessingOfMight  Buff
	BlessingOfWisdom Buff

	// Warrior
	BattleShout Buff

	// Hunter
	TrueShot        Buff
	LeaderOfThePack Buff

	// Shaman
	GraceOfAirTotem      Buff
	StrengthOfEarthTotem Buff
	ManaSpringTotem      Buff
	WindfuryTotem        Buff

	// Mage
	ArcaneBrilliance Buff

	// Druid
	GiftOfTheWild Buff

	// Shaman
	// TODO Implement bloodlust
	Bloodlust int

	// Leatherworkings
	Drums Buff

	// Priest
	PrayerOfFortitude Buff

	// Warlock
	BloodPact Buff

	// Jewelcrafting
	BraidedEterniumChain Buff
}
