package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetRing(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, rings) {
		return rings[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initRings() {
	rings = make(map[string]eq.ArmorItem)

	rings["ring of a thousand marks"] = eq.ArmorItem{
		Name: "ring of a thousand marks",
		BaseStats: eq.BaseStats{
			Stamina:     21,
			HitRating:   19,
			CritRating:  23,
			AttackPower: 44,
		},
	}

	rings["ring of the recalcitrant"] = eq.ArmorItem{
		Name: "ring of the recalcitrant",
		BaseStats: eq.BaseStats{
			Agility:     24,
			Stamina:     27,
			AttackPower: 54,
		},
	}
}
