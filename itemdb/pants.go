package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
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
}
