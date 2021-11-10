package util

import (
	"math"
	"math/rand"
)

func RollDice(chance float64) bool {
	randNumber := rand.Float64()
	if randNumber > chance {
		return false
	} else {
		return true
	}
}

func CompFloat(a float64, b float64, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
