package equipment

import (
	"reflect"
)

type Enchant BaseStats

type ArmorItem struct {
	Name string

	BaseStats

	Gems    GemSlots
	Enchant Enchant
}

type Equipment struct {
	Helm      ArmorItem
	Neck      ArmorItem
	Shoulders ArmorItem
	Cloak     ArmorItem
	Chest     ArmorItem
	Bracers   ArmorItem
	Gloves    ArmorItem
	Belt      ArmorItem
	Pants     ArmorItem
	Boots     ArmorItem

	RingOne ArmorItem
	RingTwo ArmorItem

	Ranged   RangedWeapon
	MainHand Weapon
	OffHand  Weapon

	Quiver Quiver

	TrinketOne ArmorItem
	TrinketTwo ArmorItem
}

func sumStat(name string, eq *Equipment) int {
	stat := 0
	v := reflect.ValueOf(*eq)
	for i := 0; i < v.NumField(); i++ {
		stat += int(v.Field(i).FieldByName(name).Int())
		fieldGems := v.Field(i).FieldByName("Gems").Interface()
		fieldEnchant := v.Field(i).FieldByName("Enchant").Interface()

		if v, ok := fieldGems.(GemSlots); ok {
			summedStatsValue := reflect.ValueOf(v.SummedStats())
			stat += int(summedStatsValue.FieldByName(name).Int())
		}

		if v, ok := fieldEnchant.(Enchant); ok {
			enchantValue := reflect.ValueOf(v)
			stat += int(enchantValue.FieldByName(name).Int())
		}

		fieldScope := v.Field(i).FieldByName("Scope")
		if fieldScope.IsValid() {
			if v, ok := fieldScope.Interface().(Scope); ok {
				scopeValue := reflect.ValueOf(v)
				stat += int(scopeValue.FieldByName(name).Int())
			}
		}

	}
	return stat
}

func (eq Equipment) ArmorPenetration() int {
	return sumStat("ArmorPenetration", &eq)
}

func (eq Equipment) Agility() int {
	return sumStat("Agility", &eq)
}

func (eq Equipment) Strength() int {
	return sumStat("Strength", &eq)
}

func (eq Equipment) Stamina() int {
	return sumStat("Stamina", &eq)
}

func (eq Equipment) Intellect() int {
	return sumStat("Intellect", &eq)
}

func (eq Equipment) Spirit() int {
	return sumStat("Spirit", &eq)
}

func (eq Equipment) Armor() int {
	return sumStat("Armor", &eq)
}

func (eq Equipment) AttackPower() int {
	stat := sumStat("AttackPower", &eq)

	// Check for primalstrike
	if eq.Chest.Name == "primalstrike vest" && eq.Bracers.Name == "primalstrike bracers" && eq.Belt.Name == "primalstrike belt" {
		stat += 40 //40 == primalstrike extra AP
	}

	return stat
}

func (eq Equipment) CritRating() int {
	return sumStat("CritRating", &eq)
}

func (eq Equipment) HitRating() int {
	return sumStat("HitRating", &eq)
}

func (eq Equipment) MP5() int {
	return sumStat("MP5", &eq)
}

func (eq Equipment) HasBeastLordFourSet() bool {
	// If Beast Lord
	numBeastLordItems := 0
	if eq.Chest.Name == "beast lord cuirass" {
		numBeastLordItems++
	}
	if eq.Gloves.Name == "beast lord handguards" {
		numBeastLordItems++
	}
	if eq.Helm.Name == "beast lord helm" {
		numBeastLordItems++
	}
	if eq.Pants.Name == "beast lord leggings" {
		numBeastLordItems++
	}
	if eq.Shoulders.Name == "beast lord mantle" {
		numBeastLordItems++
	}
	return numBeastLordItems >= 4
}

func (eq Equipment) HasT5FourSet() bool {
	numPieces := 0
	if eq.Chest.Name == "rift stalker hauberk" {
		numPieces++
	}
	if eq.Gloves.Name == "rift stalker gauntlets" {
		numPieces++
	}
	if eq.Helm.Name == "rift stalker helm" {
		numPieces++
	}
	if eq.Pants.Name == "rift stalker leggings" {
		numPieces++
	}
	if eq.Shoulders.Name == "rift stalker mantle" {
		numPieces++
	}
	return numPieces >= 4
}
