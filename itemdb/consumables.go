package itemdb

import (
	"fmt"
	"strings"
	cs "tbchuntersim/consumables"
	eq "tbchuntersim/equipment"
)

func GetConsumable(name string) cs.StaticConsumable {
	name = strings.ToLower(name)
	if doesItemExist(name, consumables) {
		return consumables[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func GetOil(name string) eq.Oil {
	name = strings.ToLower(name)
	if doesItemExist(name, oils) {
		return oils[name]
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}

func initConsumables() {
	consumables = make(map[string]cs.StaticConsumable)
	oils = make(map[string]eq.Oil)

	consumables["warp burger"] = cs.StaticConsumable{
		Agility:  20,
		Spirit:   20,
		Strength: 0,
	}
	consumables["spicy hot talbuk"] = cs.StaticConsumable{
		HitRating: 20,
		Spirit:    20,
	}
	consumables["kibler's bits"] = cs.StaticConsumable{
		Agility:  0,
		Strength: 20,
		Spirit:   20,
	}
	consumables["elixir of major agility"] = cs.StaticConsumable{
		Agility:    35,
		CritRating: 20,
	}
	consumables["elixir of major mageblood"] = cs.StaticConsumable{
		MP5: 16,
	}
	consumables["scroll of agility"] = cs.StaticConsumable{
		Agility: 20,
	}
	consumables["scroll of strength"] = cs.StaticConsumable{
		Strength: 20,
	}
	consumables["scroll of agility v"] = cs.StaticConsumable{
		Agility: 20,
	}
	consumables["scroll of strength v"] = cs.StaticConsumable{
		Strength: 20,
	}
	oils["adamantite weightstone"] = eq.Oil{
		BaseStats: eq.BaseStats{
			CritRating: 14,
		},
		WeaponDamage: 0,
	}
}
