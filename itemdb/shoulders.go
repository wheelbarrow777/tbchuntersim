package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetShoulder(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, shoulders) {
		return shoulders[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initShoulders() {
	shoulders = make(map[string]eq.ArmorItem)

	shoulders["beast lord mantle"] = eq.ArmorItem{
		Name: "beast lord mantle",
		BaseStats: eq.BaseStats{
			Armor:       489,
			Agility:     25,
			Intellect:   12,
			AttackPower: 34,
			MP5:         5,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.YellowGem,
				eq.RedGem,
			},
			SlottedGems: []eq.Gem{},
			Bonus: eq.BaseStats{
				Stamina: 4,
			},
		},
		Enchant: eq.Enchant{},
	}

	shoulders["rift stalker mantle"] = eq.ArmorItem{
		Name: "rift stalker mantle",
		BaseStats: eq.BaseStats{
			Armor:       700,
			Agility:     26,
			Stamina:     26,
			Intellect:   24,
			HitRating:   13,
			AttackPower: 52,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.RedGem, eq.RedGem,
			},
			Bonus: eq.BaseStats{
				Stamina: 4,
			},
		},
		Enchant: eq.Enchant{},
	}
}
