package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetBoots(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, boots) {
		return boots[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initBoots() {
	boots = make(map[string]eq.ArmorItem)

	boots["edgewalker longboots"] = eq.ArmorItem{
		Name: "edgewalker longboots",
		BaseStats: eq.BaseStats{
			Armor:       250,
			Agility:     29,
			Stamina:     28,
			HitRating:   13,
			AttackPower: 44,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem,
				eq.YellowGem,
			},
			Bonus: eq.BaseStats{
				HitRating: 3,
			},
		},
	}
}
