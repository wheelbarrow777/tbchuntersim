package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetGloves(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, gloves) {
		return gloves[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initGloves() {
	gloves = make(map[string]eq.ArmorItem)

	gloves["beast lord handguards"] = eq.ArmorItem{
		Name: "beast lord handguards",
		BaseStats: eq.BaseStats{
			Armor:       407,
			Agility:     25,
			Stamina:     12,
			Intellect:   17,
			AttackPower: 34,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem,
				eq.BlueGem,
			},
			Bonus: eq.BaseStats{
				HitRating: 3,
			},
		},
	}

	gloves["rift stalker gauntlets"] = eq.ArmorItem{
		Name: "rift stalker gauntlets",
		BaseStats: eq.BaseStats{
			Armor:       583,
			Agility:     34,
			Stamina:     29,
			Intellect:   20,
			HitRating:   19,
			AttackPower: 68,
		},
	}
}
