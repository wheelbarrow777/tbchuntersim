package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetBelt(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, belts) {
		return belts[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initBelts() {
	belts = make(map[string]eq.ArmorItem)

	belts["primalstrike belt"] = eq.ArmorItem{
		Name: "primalstrike belt",
		BaseStats: eq.BaseStats{
			Armor:       205,
			Agility:     20,
			Stamina:     32,
			AttackPower: 84,
		},
	}

	belts["belt of deep shadow"] = eq.ArmorItem{
		Name: "belt of deep shadow",
		BaseStats: eq.BaseStats{
			Armor:       227,
			Agility:     32,
			Stamina:     14,
			HitRating:   18,
			AttackPower: 66,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.BlueGem, eq.BlueGem},
			Bonus: eq.BaseStats{
				Agility: 3,
			},
		},
	}

	belts["dreghood belt of shadow protection"] = eq.ArmorItem{
		Name: "dreghood belt of shadow protection",
		BaseStats: eq.BaseStats{
			Armor:   121,
			Stamina: 31,
			Resistance: eq.Resistance{
				Shadow: 20,
			},
		},
	}

	belts["vengeance belt of shadow protection"] = eq.ArmorItem{
		Name: "vengeance belt of shadow protection",
		BaseStats: eq.BaseStats{
			Armor:   117,
			Stamina: 30,
			Resistance: eq.Resistance{
				Shadow: 19,
			},
		},
	}
}
