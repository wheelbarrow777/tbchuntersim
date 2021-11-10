package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetHelmet(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, helmets) {
		return helmets[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initHelmets() {
	helmets = make(map[string]eq.ArmorItem)

	helmets["beast lord helm"] = eq.ArmorItem{
		Name: "beast lord helm",
		BaseStats: eq.BaseStats{
			Armor:       530,
			Agility:     25,
			Stamina:     21,
			Intellect:   22,
			AttackPower: 50,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem,
				eq.MetaGem,
			},
			Bonus: eq.BaseStats{
				MP5: 2,
			},
		},
	}

	helmets["rift stalker helm"] = eq.ArmorItem{
		Name: "rift stalker helm",
		BaseStats: eq.BaseStats{
			Armor:       759,
			Agility:     40,
			Stamina:     36,
			Intellect:   25,
			AttackPower: 82,
			MP5:         10,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.YellowGem,
				eq.MetaGem,
			},
			Bonus: eq.BaseStats{
				Stamina: 6,
			},
		},
	}
}
