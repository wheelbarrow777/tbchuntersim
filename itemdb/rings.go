package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetRing(name string) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, rings) {
		return rings[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initRings() {
	rings = make(map[string]eq.ArmorItem)

	rings["ring of a thousand marks"] = eq.ArmorItem{
		Name: "ring of a thousand marks",
		BaseStats: eq.BaseStats{
			Stamina:     21,
			HitRating:   19,
			CritRating:  23,
			AttackPower: 44,
		},
	}

	rings["ring of the recalcitrant"] = eq.ArmorItem{
		Name: "ring of the recalcitrant",
		BaseStats: eq.BaseStats{
			Agility:     24,
			Stamina:     27,
			AttackPower: 54,
		},
	}

	rings["band of the ranger-general"] = eq.ArmorItem{
		Name: "band of the ranger-general",
		BaseStats: eq.BaseStats{
			Stamina:     27,
			HitRating:   18,
			CritRating:  28,
			AttackPower: 56,
		},
	}

	rings["alexandrite ring of shadow protection"] = eq.ArmorItem{
		Name: "alexandrite ring of shadow protection",
		BaseStats: eq.BaseStats{
			Stamina: 30,
			Resistance: eq.Resistance{
				Shadow: 19,
			},
		},
	}

	rings["vindicator's band of triumph"] = eq.ArmorItem{
		Name: "vindicator's band of triumph",
		BaseStats: eq.BaseStats{
			Stamina:          34,
			CritRating:       26,
			AttackPower:      44,
			ArmorPenetration: 56,
		},
	}

	rings["blue topaz band of shadow protection"] = eq.ArmorItem{
		Name: "blue topaz band of shadow protection",
		BaseStats: eq.BaseStats{
			Stamina: 26,
			Resistance: eq.Resistance{
				Shadow: 17,
			},
		},
	}

	rings["band of eternity"] = eq.ArmorItem{
		Name: "band of eternity",
		BaseStats: eq.BaseStats{
			Agility:     29,
			Stamina:     43,
			AttackPower: 60,
		},
	}

	rings["signet of primal wrath"] = eq.ArmorItem{
		Name: "signet of primal wrath",
		BaseStats: eq.BaseStats{
			Agility:          28,
			Stamina:          30,
			AttackPower:      58,
			ArmorPenetration: 126,
		},
	}

	rings["band of the eternal champion"] = eq.ArmorItem{
		Name: "band of the eternal champion",
		BaseStats: eq.BaseStats{
			Agility:     29,
			Stamina:     43,
			AttackPower: 60,
		},
	}
}
