package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
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

	bracers["nightfall wristguards"] = eq.ArmorItem{
		Name: "nightfall wristguards",
		BaseStats: eq.BaseStats{
			Armor:       153,
			Agility:     24,
			Stamina:     22,
			AttackPower: 46,
		},
	}

	bracers["expedition bracers of shadow protection"] = eq.ArmorItem{
		Name: "expedition bracers of shadow protection",
		BaseStats: eq.BaseStats{
			Armor:       118,
			Stamina:     31,
			AttackPower: 24,
			Resistance: eq.Resistance{
				Shadow: 20,
			},
		},
	}

	bracers["skettis bracer of shadow protection"] = eq.ArmorItem{
		Name: "skettis bracer of shadow protection",
		BaseStats: eq.BaseStats{
			Armor:   251,
			Stamina: 29,
			Resistance: eq.Resistance{
				Shadow: 19,
			},
		},
	}

	bracers["insidious bands"] = eq.ArmorItem{
		Name: "insidious bands",
		BaseStats: eq.BaseStats{
			Armor:       194,
			Agility:     28,
			Stamina:     28,
			HitRating:   12,
			AttackPower: 58,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.YellowGem},
			Bonus: eq.BaseStats{
				Agility: 2,
			},
		},
	}
}
