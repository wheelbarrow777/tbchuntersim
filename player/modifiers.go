package player

import (
	"math"
	"reflect"
)

type timerModifier struct {
	RapidFire  float64
	TBW        float64
	QuickShots float64
	Bloodlust  float64
	Racial     float64

	HastePotion float64
	Drums       float64

	BeastLordArmorPen float64

	// Trinkets
	DST             float64
	BloodlustBrooch float64
}

type ActiveModifiers struct {
	TimerModifiers timerModifier
}

const (
	HASTE_RATING_RATIO    = 15.77
	HUNTERS_MARK_AP_BONUS = 440

	BEAST_LORD_ARMOR_IGNORE = 600
	DST_HASTE_BONUS         = 325
	HASTE_POTION_BONUS      = 400

	CRIT_RATING_RATIO = 22.0765
	HIT_RATING_RATION = 15.76

	AGILITY_TO_CRIT = 40.0

	DRUMS_OF_BATTLE_HASTE_BONUS = 80
)

func (am *ActiveModifiers) ReduceModifierTime(duration float64) {
	v := reflect.ValueOf(&am.TimerModifiers)
	for i := 0; i < v.Elem().NumField(); i++ {
		currentValue := v.Elem().Field(i).Float()
		newValue := math.Max(0, currentValue-duration)
		v.Elem().Field(i).SetFloat(newValue)
	}
}
