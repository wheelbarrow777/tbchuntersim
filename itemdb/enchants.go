package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func GetEnchant(name string) eq.Enchant {
	name = strings.ToLower(name)
	if doesItemExist(name, enchants) {
		return enchants[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func GetScope(name string) eq.Scope {
	name = strings.ToLower(name)
	if doesItemExist(name, scopes) {
		return scopes[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initEnchants() {
	enchants = make(map[string]eq.Enchant)
	scopes = make(map[string]eq.Scope)

	enchants["glyph of ferocity"] = eq.Enchant{
		AttackPower: 34,
		HitRating:   16,
	}

	enchants["greater inscription of vengeance"] = eq.Enchant{
		AttackPower: 30,
		CritRating:  10,
	}

	enchants["enchant cloak - greater agility"] = eq.Enchant{
		Agility: 12,
	}

	enchants["enchant chest - exceptional stats"] = eq.Enchant{
		Agility:   6,
		Stamina:   6,
		Intellect: 6,
		Spirit:    6,
		Strength:  6,
	}

	enchants["enchant bracer - assault"] = eq.Enchant{
		AttackPower: 24,
	}

	enchants["enchant gloves - superior agility"] = eq.Enchant{
		Agility: 15,
	}

	enchants["nethercobra leg armor"] = eq.Enchant{
		AttackPower: 50,
		CritRating:  12,
	}

	enchants["enchant boots - dexterity"] = eq.Enchant{
		Agility: 12,
	}

	enchants["enchant weapon - agility"] = eq.Enchant{
		Agility: 15,
	}

	enchants["glyph of shadow warding"] = eq.Enchant{
		Resistance: eq.Resistance{
			Shadow: 20,
		},
	}

	enchants["shadow armor kit"] = eq.Enchant{
		Resistance: eq.Resistance{
			Shadow: 8,
		},
	}

	enchants["enchant cloak - greater shadow resistance"] = eq.Enchant{
		Resistance: eq.Resistance{
			Shadow: 15,
		},
	}

	scopes["stabilized eternium scope"] = eq.Scope{
		Damage: 0,
		BaseStats: eq.BaseStats{
			CritRating: 28,
		},
	}
}
