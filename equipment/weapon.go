package equipment

import "math/rand"

type Weapon struct {
	DamageMin   int
	DamageMax   int
	Speed       float64
	IsTwoHanded bool
	ArmorItem
	Oil Oil
}

type Oil struct {
	BaseStats
	WeaponDamage int
}

type Scope struct {
	Damage int
	BaseStats
}

type RangedWeapon struct {
	Weapon
	Scope   Scope
	AmmoDPS float64
}

func (w RangedWeapon) AmmoDamage() float64 {
	return w.Speed * w.AmmoDPS
}

func (w RangedWeapon) AverageDamage() float64 {
	return float64(w.DamageMin + rand.Intn(w.DamageMax-w.DamageMin))
}
