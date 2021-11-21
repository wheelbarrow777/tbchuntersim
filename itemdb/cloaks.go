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
}
