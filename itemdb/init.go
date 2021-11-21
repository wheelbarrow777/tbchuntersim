package itemdb

import (
	"reflect"
	cs "tbchuntersim/consumables"
	eq "tbchuntersim/equipment"
)

var (
	helmets       map[string]eq.ArmorItem
	neck          map[string]eq.ArmorItem
	shoulders     map[string]eq.ArmorItem
	cloaks        map[string]eq.ArmorItem
	chests        map[string]eq.ArmorItem
	bracers       map[string]eq.ArmorItem
	gloves        map[string]eq.ArmorItem
	belts         map[string]eq.ArmorItem
	pants         map[string]eq.ArmorItem
	boots         map[string]eq.ArmorItem
	rings         map[string]eq.ArmorItem
	trinkets      map[string]eq.ArmorItem
	meleeWeapons  map[string]eq.Weapon
	rangedWeapons map[string]eq.RangedWeapon
	consumables   map[string]cs.StaticConsumable
	oils          map[string]eq.Oil
	gems          map[string]eq.Gem
	enchants      map[string]eq.Enchant
	scopes        map[string]eq.Scope
	quivers       map[string]eq.Quiver
)

func doesItemExist(name string, m interface{}) bool {
	v := reflect.ValueOf(m).MapRange()
	for v.Next() {
		if v.Key().String() == name {
			return true
		}
	}

	return false
}

func init() {
	initHelmets()
	initNecklaces()
	initShoulders()
	initCloaks()
	initChests()
	initBracers()
	initGloves()
	initBelts()
	initPants()
	initBoots()
	initRings()
	initTrinkets()
	initWeapons()
	initConsumables()
	initGems()
	initEnchants()
}
