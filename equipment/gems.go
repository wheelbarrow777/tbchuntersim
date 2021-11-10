package equipment

import log "github.com/sirupsen/logrus"

type GemColor string

type Gem struct {
	Name  string
	Color GemColor
	BaseStats
}

type GemSlots struct {
	SlotColors  []GemColor
	SlottedGems []Gem
	Bonus       BaseStats
}

func (gs *GemSlots) AddGem(g Gem) {
	gs.SlottedGems = append(gs.SlottedGems, g)
}

func (gs GemSlots) bonusAchieved() bool {
	// How many of each gem do we need?
	gemsRequired := make(map[GemColor]int)
	availableHybridgems := []Gem{}
	for _, slotColor := range gs.SlotColors {
		gemsRequired[slotColor]++
	}

	// For each of the gems required, check!
	for _, gem := range gs.SlottedGems {
		if gem.Color.isHybrid() {
			availableHybridgems = append(availableHybridgems, gem)
			continue
		}
		gemsRequired[gem.Color]--
	}

	// The map must be zero for all values
	for gem, value := range gemsRequired {
		if value != 0 {
			// Depending on the color, check if we can use one of the hybrid gems
			hybridGemCovered := false
			for i, hGem := range availableHybridgems {
				if gem.validSlot(hGem.Color) {
					// Remove the gem as it's used
					availableHybridgems[i] = availableHybridgems[len(availableHybridgems)-1]
					availableHybridgems = availableHybridgems[:len(availableHybridgems)-1]
					hybridGemCovered = true
					break
				}
			}
			// If no hybrid gem was found to cover, the bonus is not achieved
			if !hybridGemCovered {
				return false
			}
		}
	}

	return true
}

func (gs GemSlots) SummedStats() BaseStats {
	base := BaseStats{}
	for _, gem := range gs.SlottedGems {
		base.Add(gem.BaseStats)
	}

	if gs.bonusAchieved() {
		base.Add(gs.Bonus)
	}

	return base
}

const (
	RedGem    = "red"
	BlueGem   = "blue"
	YellowGem = "yellow"
	MetaGem   = "meta"
	Purple    = "purple"
	Orange    = "orange"
)

func (first GemColor) validSlot(second GemColor) bool {
	if first == RedGem {
		if second == RedGem || second == Orange || second == Purple {
			return true
		} else {
			return false
		}
	}

	if first == BlueGem {
		if second == BlueGem || second == Purple {
			return true
		} else {
			return false
		}
	}

	if first == YellowGem {
		if second == YellowGem || second == Orange {
			return true
		} else {
			return false
		}
	}

	if first == MetaGem {
		return second == MetaGem
	}

	log.WithFields(log.Fields{
		"FirstGem":  first,
		"SecondGem": second,
	}).Warn("Attempted to check if a non-main gem fit in a slot")
	return false
}

func (g GemColor) isHybrid() bool {
	if g == Purple || g == Orange {
		return true
	} else {
		return false
	}
}
