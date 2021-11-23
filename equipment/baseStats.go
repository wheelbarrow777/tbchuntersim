package equipment

type Resistance struct {
	Frost  int
	Shadow int
}

type BaseStats struct {
	Armor int

	Agility   int
	Stamina   int
	Intellect int
	Spirit    int
	Strength  int

	AttackPower      int
	CritRating       int
	HasteRating      int
	ArmorPenetration int

	Resistance Resistance

	MP5       int
	HitRating int
}

func (first *BaseStats) Add(second BaseStats) {
	first.Armor += second.Armor
	first.Agility += second.Agility
	first.Stamina += second.Stamina
	first.Intellect += second.Intellect
	first.Spirit += second.Spirit
	first.Strength += second.Strength

	first.AttackPower += second.AttackPower
	first.CritRating += second.CritRating
	first.HasteRating += second.CritRating
	first.HasteRating += second.HasteRating
	first.ArmorPenetration += second.ArmorPenetration

	first.MP5 += second.MP5
	first.HitRating += second.HitRating
}
