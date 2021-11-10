package player

// TODO: Memoize
func (p Player) BaseHaste() float64 {
	serp := 1.0 + float64(p.Talents.BM.SerpentsSwiftness)*0.04
	quiver := p.Equipment.Quiver.Speed

	// TODO: Add gear modifier based on gear
	gearModifier := 1.0

	return serp * quiver * gearModifier
}

func (p Player) BonusHaste() float64 {
	rating := 1.0
	if p.Am.TimerModifiers.RapidFire > 0 {
		rating = rating * 1.4
	}

	if p.Am.TimerModifiers.QuickShots > 0 {
		rating = rating * 1.15
	}

	if p.Am.TimerModifiers.Bloodlust > 0 {
		rating = rating * 1.3
	}

	if p.Race.Name == "troll" {
		if p.Am.TimerModifiers.Racial > 0 {
			rating = rating * (1 + 0.12)
		}
	}

	return rating
}

func (p Player) BonusHasteRating() float64 {
	rating := p.hasteRating

	if p.Am.TimerModifiers.DST > 0 {
		rating += DST_HASTE_BONUS
	}

	if p.Am.TimerModifiers.HastePotion > 0 {
		rating += HASTE_POTION_BONUS
	}

	return rating
}

func (p Player) TotalHaste() float64 {
	h := p.BaseHaste() * p.BonusHaste() * (1 + p.BonusHasteRating()/HASTE_RATING_RATIO/100)
	return h
}

func (p Player) ArmorPenetrationRatio() float64 {
	eq := p.Equipment

	amp := eq.ArmorPenetration()

	if p.Am.BeastLordArmorPen > 0 {
		amp += BEAST_LORD_ARMOR_IGNORE
	}

	return float64(amp)
}

func (p Player) RealSpeed() float64 {
	return p.Equipment.Ranged.Speed / p.TotalHaste()
}
