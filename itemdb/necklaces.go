package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetNecklace(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, neck) {
		return neck[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initNecklaces() {
	neck = make(map[string]eq.ArmorItem)

	neck["choker of vile intent"] = eq.ArmorItem{
		Name: "choker of vile intent",
		BaseStats: eq.BaseStats{
			Agility:     20,
			Stamina:     18,
			AttackPower: 42,
			HitRating:   18,
		},
	}

	neck["telonicus's pendant of mayhem"] = eq.ArmorItem{
		Name: "telonicus's pendant of mayhem",
		BaseStats: eq.BaseStats{
			Agility:     27,
			Stamina:     26,
			AttackPower: 70,
		},
	}

	neck["medallion of karabor"] = eq.ArmorItem{
		Name: "medallion of karabor",
		BaseStats: eq.BaseStats{
			Stamina: 49,
			Resistance: eq.Resistance{
				Shadow: 40,
			},
		},
	}

	neck["choker of endless nightmares"] = eq.ArmorItem{
		Name: "Choker of Endless Nightmares",
		BaseStats: eq.BaseStats{
			HasteRating: 21,
			CritRating:  27,
			AttackPower: 72,
		},
	}

	neck["blessed medallion of karabor"] = eq.ArmorItem{
		Name: "blessed medallion of karabor",
		BaseStats: eq.BaseStats{
			Stamina: 49,
			Resistance: eq.Resistance{
				Shadow: 40,
			},
		},
	}
}
