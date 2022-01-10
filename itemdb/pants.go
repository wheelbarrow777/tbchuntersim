package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetPants(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, pants) {
		return pants[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initPants() {
	pants = make(map[string]eq.ArmorItem)

	pants["beast lord leggings"] = eq.ArmorItem{
		Name: "beast lord leggings",
		BaseStats: eq.BaseStats{
			Armor:       570,
			Agility:     30,
			Stamina:     25,
			Intellect:   19,
			AttackPower: 52,
			MP5:         7,
		},
	}

	pants["void reaver greaves"] = eq.ArmorItem{
		Name: "void reaver greaves",
		BaseStats: eq.BaseStats{
			Armor:       787,
			Agility:     37,
			Stamina:     33,
			Intellect:   24,
			AttackPower: 88,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.RedGem, eq.YellowGem, eq.BlueGem},
			Bonus: eq.BaseStats{
				Agility: 4,
			},
		},
	}

	pants["leggings of murderous intent"] = eq.ArmorItem{
		Name: "leggings of murderous intent",
		BaseStats: eq.BaseStats{
			Armor:       380,
			Agility:     45,
			Stamina:     31,
			CritRating:  37,
			AttackPower: 92,
		},
	}

	pants["midnight legguards"] = eq.ArmorItem{
		Name: "midnight legguards",
		BaseStats: eq.BaseStats{
			Armor:       305,
			Stamina:     30,
			HitRating:   17,
			CritRating:  27,
			AttackPower: 64,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.RedGem, eq.RedGem, eq.BlueGem},
			Bonus: eq.BaseStats{
				HitRating: 4,
			},
		},
	}

	pants["rift stalker leggings"] = eq.ArmorItem{
		Name: "rift stalker leggings",
		BaseStats: eq.BaseStats{
			Armor:       817,
			Agility:     40,
			Stamina:     39,
			Intellect:   26,
			HitRating:   18,
			AttackPower: 92,
			MP5:         7,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.RedGem},
			Bonus: eq.BaseStats{
				CritRating: 2,
			},
		},
	}
}
