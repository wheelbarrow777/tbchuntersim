package itemdb

import (
	"fmt"
	eq "huntsim/equipment"
	"strings"
)

func GetGem(name string) eq.Gem {
	name = strings.ToLower(name)
	if doesItemExist(name, gems) {
		return gems[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initGems() {
	gems = make(map[string]eq.Gem)

	gems["delicate living ruby"] = eq.Gem{
		Name:  "delicate living ruby",
		Color: eq.RedGem,
		BaseStats: eq.BaseStats{
			Agility: 8,
		},
	}

	gems["relentless earthstorm diamond"] = eq.Gem{
		Name:  "relentless earthstorm diamond",
		Color: eq.MetaGem,
		BaseStats: eq.BaseStats{
			Agility: 12,
		},
	}

	gems["shifting nightseye"] = eq.Gem{
		Name:  "shifting nightseye",
		Color: eq.Purple,
		BaseStats: eq.BaseStats{
			Agility: 4,
			Stamina: 6,
		},
	}

	gems["wicked noble topaz"] = eq.Gem{
		Name:  "wicked noble topaz",
		Color: eq.Orange,
		BaseStats: eq.BaseStats{
			CritRating:  4,
			AttackPower: 8,
		},
	}

	gems["stone of blades"] = eq.Gem{
		Name:  "stone of blades",
		Color: eq.YellowGem,
		BaseStats: eq.BaseStats{
			CritRating: 12,
		},
	}

}
