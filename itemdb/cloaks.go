package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetCloak(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, cloaks) {
		return cloaks[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initCloaks() {
	cloaks = make(map[string]eq.ArmorItem)

	cloaks["cloak of the pit stalker"] = eq.ArmorItem{
		Name: "cloak of the pit stalker",
		BaseStats: eq.BaseStats{
			Armor:       105,
			Stamina:     28,
			CritRating:  24,
			AttackPower: 56,
		},
	}

	cloaks["thalassian wildercloak"] = eq.ArmorItem{
		Name: "thalassian wildercloak",
		BaseStats: eq.BaseStats{
			Armor:       116,
			Agility:     28,
			Stamina:     27,
			AttackPower: 68,
		},
	}

	cloaks["abyssal shroud"] = eq.ArmorItem{
		Name: "abyssal shroud",
		BaseStats: eq.BaseStats{
			Stamina: 15,
			Resistance: eq.Resistance{
				Shadow: 18,
			},
		},
	}

	cloaks["vengeance wrap"] = eq.ArmorItem{
		Name: "vengeance wrap",
		BaseStats: eq.BaseStats{
			Armor:       89,
			CritRating:  23,
			AttackPower: 52,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.RedGem},
			Bonus: eq.BaseStats{
				HitRating: 2,
			},
		},
	}

	cloaks["forest shroud of shadow protection"] = eq.ArmorItem{
		Name: "forest shroud of shadow protection",
		BaseStats: eq.BaseStats{
			Stamina: 27,
			Resistance: eq.Resistance{
				Shadow: 17,
			},
		},
	}
}
