package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetMeleeWeapon(name string) eq.Weapon {
	name = strings.ToLower(name)
	if doesItemExist(name, meleeWeapons) {
		return meleeWeapons[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func GetRangedWeapon(name string) eq.RangedWeapon {
	name = strings.ToLower(name)
	if doesItemExist(name, rangedWeapons) {
		return rangedWeapons[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func GetQuiver(name string) eq.Quiver {
	name = strings.ToLower(name)
	if doesItemExist(name, quivers) {
		return quivers[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initWeapons() {
	meleeWeapons = make(map[string]eq.Weapon)
	rangedWeapons = make(map[string]eq.RangedWeapon)
	quivers = make(map[string]eq.Quiver)

	meleeWeapons["stellaris"] = eq.Weapon{
		ArmorItem: eq.ArmorItem{
			Name: "stellaris",
			BaseStats: eq.BaseStats{
				Agility:     21,
				Stamina:     12,
				AttackPower: 22,
			},
		},
		DamageMin:   95,
		DamageMax:   177,
		Speed:       1.9,
		IsTwoHanded: false,
	}

	meleeWeapons["stormreaver warblades"] = eq.Weapon{
		ArmorItem: eq.ArmorItem{
			Name: "stormreave warblades",
			BaseStats: eq.BaseStats{
				Stamina:     13,
				CritRating:  21,
				AttackPower: 22,
			},
		},
		DamageMin:   80,
		DamageMax:   149,
		Speed:       1.6,
		IsTwoHanded: false,
	}

	rangedWeapons["sunfury bow of the phoenix"] = eq.RangedWeapon{
		Weapon: eq.Weapon{
			ArmorItem: eq.ArmorItem{
				Name: "sunfury bow of the phoenix",
				BaseStats: eq.BaseStats{
					Agility:     19,
					AttackPower: 34,
				},
			},
			DamageMin: 169,
			DamageMax: 314,
			Speed:     2.9,
		},
		AmmoDPS: 43,
	}

	quivers["clefthoof hide quiver"] = eq.Quiver{
		Speed: 1.15,
	}
}
