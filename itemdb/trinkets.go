package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetTrinket(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, trinkets) {
		return trinkets[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initTrinkets() {
	trinkets = make(map[string]eq.ArmorItem)

	trinkets["bloodlust brooch"] = eq.ArmorItem{
		Name: "bloodlust brooch",
		BaseStats: eq.BaseStats{
			AttackPower: 72,
		},
	}
	trinkets["dragonspine trophy"] = eq.ArmorItem{
		Name: "dragonspine trophy",
		BaseStats: eq.BaseStats{
			AttackPower: 40,
		},
	}
	trinkets["ultra-flash shadow reflector"] = eq.ArmorItem{
		Name: "ultra-flash shadow reflector",
		BaseStats: eq.BaseStats{
			Resistance: eq.Resistance{
				Shadow: 20,
			},
		},
	}

	trinkets["gnomish battle chicken"] = eq.ArmorItem{
		Name:      "gnomish battle chicken",
		BaseStats: eq.BaseStats{},
	}

	trinkets["madness of the betrayer"] = eq.ArmorItem{
		Name: "madness of the betrayer",
		BaseStats: eq.BaseStats{
			HitRating:   20,
			AttackPower: 84,
		},
	}

	trinkets["berserker's call"] = eq.ArmorItem{
		Name: "berserker's call",
		BaseStats: eq.BaseStats{
			AttackPower: 90,
		},
	}
}
