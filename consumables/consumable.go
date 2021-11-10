package consumables

import (
	"huntsim/equipment"
)

type StaticConsumable equipment.BaseStats

type Player []StaticConsumable

type Pet []StaticConsumable

type StaticConsumables struct {
	Player Player
	Pet    Pet
}

func (sc StaticConsumables) SummedStatsPlayer() equipment.BaseStats {
	b := equipment.BaseStats{}
	for _, c := range sc.Player {
		b.Armor += c.Armor
		b.Agility += c.Agility
		b.Stamina += c.Stamina
		b.Intellect += c.Intellect
		b.Spirit += c.Spirit
		b.Strength += c.Strength
		b.AttackPower += c.AttackPower
		b.CritRating += c.CritRating
		b.ArmorPenetration += c.ArmorPenetration
		b.MP5 += c.MP5
		b.HitRating += c.HitRating
	}

	return b
}
