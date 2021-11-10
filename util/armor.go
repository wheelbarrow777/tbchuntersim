package util

func CalculateReducedArmorDamage(damage float64, armor float64) float64 {
	// TODO: Make moblvl dyanmic
	moblvl := 70.0
	reduction := armor / (armor - 22167.5 + (467.5 * moblvl))
	return damage * (1 - reduction)
}
