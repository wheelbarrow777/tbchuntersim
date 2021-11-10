package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetChest(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, chests) {
		return chests[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initChests() {
	chests = make(map[string]eq.ArmorItem)

	chests["primalstrike vest"] = eq.ArmorItem{
		Name: "primalstrike vest",
		BaseStats: eq.BaseStats{
			Armor:       379,
			Agility:     38,
			Stamina:     39,
			HitRating:   12,
			AttackPower: 108,
		},
	}

	chests["rift stalker hauberk"] = eq.ArmorItem{
		Name: "rift stalker hauberk",
		BaseStats: eq.BaseStats{
			Armor:       934,
			Agility:     40,
			Stamina:     40,
			Intellect:   19,
			HitRating:   19,
			AttackPower: 80,
			MP5:         7,
		},
		Gems: eq.GemSlots{
			SlotColors: []eq.GemColor{
				eq.BlueGem, eq.YellowGem, eq.YellowGem,
			},
			Bonus: eq.BaseStats{
				Agility: 4,
			},
		},
	}
}
