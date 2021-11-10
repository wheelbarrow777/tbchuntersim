package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetBracers(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, bracers) {
		return bracers[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initBracers() {
	bracers = make(map[string]eq.ArmorItem)

	bracers["primalstrike bracers"] = eq.ArmorItem{
		Name: "primalstrike bracers",
		BaseStats: eq.BaseStats{
			Armor:       159,
			Agility:     15,
			Stamina:     21,
			AttackPower: 64,
		},
	}

	bracers["vambraces of ending"] = eq.ArmorItem{
		Name: "vambraces of ending",
		BaseStats: eq.BaseStats{
			Agility:     24,
			Stamina:     24,
			AttackPower: 52,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.BlueGem,
			},
			Bonus: eq.BaseStats{
				AttackPower: 4,
			},
		},
	}
}
