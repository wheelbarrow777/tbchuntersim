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
}
