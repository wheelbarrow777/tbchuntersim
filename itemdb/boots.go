package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
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

	boots["cobra-lash boots"] = eq.ArmorItem{
		Name: "cobra-lash boots",
		BaseStats: eq.BaseStats{
			Armor:       665,
			Agility:     33,
			Stamina:     25,
			Intellect:   25,
			AttackPower: 66,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.BlueGem,
				eq.RedGem,
			},
			Bonus: eq.BaseStats{
				Agility: 3,
			},
		},
	}

	boots["frostsaber boots"] = eq.ArmorItem{
		Name: "frostsaber boots",
		BaseStats: eq.BaseStats{
			Armor: 99,
			Resistance: eq.Resistance{
				Shadow: 12,
				Frost:  12,
			},
		},
	}

	boots["fel leather boots"] = eq.ArmorItem{
		Name: "fel leather boots",
		BaseStats: eq.BaseStats{
			Armor:       196,
			HitRating:   25,
			CritRating:  17,
			AttackPower: 36,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.YellowGem, eq.RedGem},
			Bonus: eq.BaseStats{
				AttackPower: 6,
			},
		},
	}

	boots["murkblood boots of shadow protection"] = eq.ArmorItem{
		Name: "murkblood boots of shadow protection",
		BaseStats: eq.BaseStats{
			Armor:   181,
			Stamina: 40,
			Resistance: eq.Resistance{
				Shadow: 26,
			},
		},
	}

	boots["rocket boots xtreme"] = eq.ArmorItem{
		Name: "rocket boots xtreme",
		BaseStats: eq.BaseStats{
			Armor:       196,
			AttackPower: 80,
			Resistance: eq.Resistance{
				Shadow: 8,
			},
		},
	}
}
