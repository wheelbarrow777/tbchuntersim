package util

import log "github.com/sirupsen/logrus"

var lowManaWarningGive = false

func NotifyLowMana() {
	if lowManaWarningGive {
		return
	}

	log.Warn("Low Mana!")
	lowManaWarningGive = true
}
