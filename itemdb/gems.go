package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
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

	gems["shifting tanzanite"] = eq.Gem{
		Name:  "shifting tanzanite",
		Color: eq.Purple,
		BaseStats: eq.BaseStats{
			Agility: 5,
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

	gems["bold ornate ruby"] = eq.Gem{
		Name:  "bold ornate ruby",
		Color: eq.RedGem,
		BaseStats: eq.BaseStats{
			AttackPower: 20,
		},
	}

	gems["wicked pyrestone"] = eq.Gem{
		Name:  "wicked pyrestone",
		Color: eq.Orange,
		BaseStats: eq.BaseStats{
			CritRating:  5,
			AttackPower: 10,
		},
	}

	gems["smooth lionseye"] = eq.Gem{
		Name:  "smooth lionseye",
		Color: eq.YellowGem,
		BaseStats: eq.BaseStats{
			CritRating: 10,
		},
	}

	gems["glinting fire opal"] = eq.Gem{
		Name:  "glinting fire opal",
		Color: eq.Orange,
		BaseStats: eq.BaseStats{
			Agility:   5,
			HitRating: 4,
		},
	}

	gems["inscribed ornate topaz"] = eq.Gem{
		Name:  "inscribed ornate topaz",
		Color: eq.Orange,
		BaseStats: eq.BaseStats{
			AttackPower: 10,
			CritRating:  5,
		},
	}

	gems["void sphere"] = eq.Gem{
		Name:  "void sphere",
		Color: eq.Prismatic,
		BaseStats: eq.BaseStats{
			Resistance: eq.Resistance{
				Frost:  4,
				Shadow: 4,
				Fire:   4,
				Nature: 4,
				Arcane: 4,
			},
		},
	}

}
