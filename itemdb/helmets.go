package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
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

	helmets["cogspinner goggles of shadow"] = eq.ArmorItem{
		Name: "cogspinner goggles of shadow",
		BaseStats: eq.BaseStats{
			Armor:   220,
			Stamina: 59,
			Resistance: eq.Resistance{
				Shadow: 20,
			},
		},
	}

	helmets["cursed vision of sargeras"] = eq.ArmorItem{
		Name: "cursed vision of sargeras",
		BaseStats: eq.BaseStats{
			Armor:       385,
			Agility:     39,
			Stamina:     46,
			HitRating:   21,
			CritRating:  38,
			AttackPower: 108,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.MetaGem,
				eq.YellowGem,
			},
			Bonus: eq.BaseStats{
				Stamina: 6,
			},
		},
	}

	helmets["gronnstalker's helmet"] = eq.ArmorItem{
		Name: "gronnstalker's helmet",
		BaseStats: eq.BaseStats{
			Armor:       830,
			Agility:     45,
			Stamina:     45,
			Intellect:   29,
			AttackPower: 90,
			MP5:         8,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem,
				eq.MetaGem,
			},
			Bonus: eq.BaseStats{
				CritRating: 4,
			},
		},
	}

	helmets["vengeful gladiator's chain helm"] = eq.ArmorItem{
		Name: "vengeful gladiator's chain helm",
		BaseStats: eq.BaseStats{
			Armor:            830,
			Agility:          33,
			Stamina:          61,
			Intellect:        22,
			HitRating:        12,
			CritRating:       26,
			AttackPower:      58,
			ArmorPenetration: 84,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{eq.MetaGem, eq.RedGem},
		},
	}
}
