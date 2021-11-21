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
}
