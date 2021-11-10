package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
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
}
