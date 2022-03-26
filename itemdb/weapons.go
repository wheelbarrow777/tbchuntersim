package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
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

	meleeWeapons["talon of the phoenix"] = eq.Weapon{
		DamageMin:   182,
		DamageMax:   339,
		Speed:       2.7,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "talon of the phoenix",
			BaseStats: eq.BaseStats{
				HitRating:   15,
				CritRating:  19,
				AttackPower: 52,
			},
		},
	}

	meleeWeapons["claw of the phoenix"] = eq.Weapon{
		DamageMin:   101,
		DamageMax:   189,
		Speed:       1.5,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "claw of the phoenix",
			BaseStats: eq.BaseStats{
				Agility:     21,
				Stamina:     30,
				AttackPower: 40,
			},
		},
	}

	meleeWeapons["claw of the watcher"] = eq.Weapon{
		DamageMin:   125,
		DamageMax:   233,
		Speed:       2.5,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "claw of the watcher",
			BaseStats: eq.BaseStats{
				CritRating:  12,
				AttackPower: 24,
			},
			Gems: eq.GemSlots{
				SlotColors: []eq.GemColor{
					eq.RedGem, eq.BlueGem,
				},
				Bonus: eq.BaseStats{
					CritRating: 3,
				},
			},
		},
	}

	meleeWeapons["blade of the unrequited"] = eq.Weapon{
		DamageMin:   112,
		DamageMax:   168,
		Speed:       1.6,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "blade of the unrequited",
			BaseStats: eq.BaseStats{
				Stamina:     13,
				CritRating:  9,
				AttackPower: 18,
			},
			Gems: eq.GemSlots{
				SlotColors: []eq.GemColor{eq.RedGem, eq.YellowGem, eq.BlueGem},
				Bonus: eq.BaseStats{
					AttackPower: 8,
				},
			},
		},
	}

	meleeWeapons["twinblade of the phoenix"] = eq.Weapon{
		DamageMin:   375,
		DamageMax:   564,
		Speed:       3.6,
		IsTwoHanded: true,
		ArmorItem: eq.ArmorItem{
			Name: "twinblade of the phoenix",
			BaseStats: eq.BaseStats{
				Stamina:     53,
				CritRating:  37,
				AttackPower: 110,
			},
			Gems: eq.GemSlots{
				SlotColors: []eq.GemColor{
					eq.RedGem, eq.RedGem, eq.RedGem,
				},
				Bonus: eq.BaseStats{
					AttackPower: 8,
				},
			},
		},
	}

	meleeWeapons["boundless agony"] = eq.Weapon{
		DamageMin:   144,
		DamageMax:   217,
		Speed:       1.8,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "boundless agony",
			BaseStats: eq.BaseStats{
				CritRating:       24,
				ArmorPenetration: 210,
			},
		},
	}

	meleeWeapons["tracker's blade"] = eq.Weapon{
		DamageMin:   105,
		DamageMax:   196,
		Speed:       1.5,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "tracker's blade",
			BaseStats: eq.BaseStats{
				HitRating:   20,
				CritRating:  23,
				AttackPower: 44,
			},
		},
	}

	meleeWeapons["dagger of bad mojo"] = eq.Weapon{
		DamageMin:   137,
		DamageMax:   207,
		Speed:       1.8,
		IsTwoHanded: false,
		ArmorItem: eq.ArmorItem{
			Name: "dagger of bad mojo",
			BaseStats: eq.BaseStats{
				Agility:          21,
				AttackPower:      40,
				ArmorPenetration: 140,
			},
		},
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

	rangedWeapons["serpent spine longbow"] = eq.RangedWeapon{
		Weapon: eq.Weapon{
			DamageMin:   217,
			DamageMax:   327,
			Speed:       3.0,
			IsTwoHanded: false,
			ArmorItem: eq.ArmorItem{
				Name: "serpent spine longbow",
				BaseStats: eq.BaseStats{
					Stamina:     17,
					CritRating:  16,
					AttackPower: 38,
				},
			},
		},
	}

	rangedWeapons["bristleblitz striker"] = eq.RangedWeapon{
		Weapon: eq.Weapon{
			DamageMin:   201,
			DamageMax:   374,
			Speed:       3.0,
			IsTwoHanded: false,
			ArmorItem: eq.ArmorItem{
				Name: "bristleblitz striker",
				BaseStats: eq.BaseStats{
					Stamina:    28,
					CritRating: 25,
				},
			},
		},
	}

	quivers["clefthoof hide quiver"] = eq.Quiver{
		Speed: 1.15,
	}

	quivers["quiver of a thousand feathers"] = eq.Quiver{
		Speed: 1.5,
	}
}
