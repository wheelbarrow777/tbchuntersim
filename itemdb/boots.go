package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetBoots(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, boots) {
		return boots[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initBoots() {
	boots = make(map[string]eq.ArmorItem)

	boots["edgewalker longboots"] = eq.ArmorItem{
		Name: "edgewalker longboots",
		BaseStats: eq.BaseStats{
			Armor:       250,
			Agility:     29,
			Stamina:     28,
			HitRating:   13,
			AttackPower: 44,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem,
				eq.YellowGem,
			},
			Bonus: eq.BaseStats{
				HitRating: 3,
			},
		},
	}

	boots["cobra-lash boots"] = eq.ArmorItem{
		Name: "cobra-lash boots",
		BaseStats: eq.BaseStats{
			Armor:       665,
			Agility:     33,
			Stamina:     25,
			Intellect:   25,
			AttackPower: 66,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.BlueGem,
				eq.RedGem,
			},
			Bonus: eq.BaseStats{
				Agility: 3,
			},
		},
	}
}
