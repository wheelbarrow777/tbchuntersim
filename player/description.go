package player

import (
	"fmt"
	"io"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func (p Player) renderStatsCard(t io.Writer) {
	statTable := tablewriter.NewWriter(t)
	statTable.SetHeader([]string{"Stat", "Value"})
	statData := [][]string{
		{"Agility", fmt.Sprintf("%d", p.agility)},
		{"Strength", fmt.Sprintf("%d", p.strength)},
		{"Stamina", fmt.Sprintf("%d", p.stamina)},
		{"Intellect", fmt.Sprintf("%d", p.intellect)},
		{"Spirit", fmt.Sprintf("%d", p.spirit)},

		{"Melee AP", fmt.Sprintf("%d", p.meleeAttackPower)},
		{"Ranged AP", fmt.Sprintf("%d", p.rangedAttackPower)},
		{"Crit Rating", fmt.Sprintf("%d", p.critRating)},
		{"Hit Rating", fmt.Sprintf("%d", p.hitRating)},

		{"MP5", fmt.Sprintf("%d", p.mp5)},
	}

	// TODO: Add Print
	// if viper.GetBool("print-resistance") {
	// 	statData = append(statData, []string{
	// 		"Shadow Resistance", fmt.Sprintf("%d", p.Equipment.Helm.Gems.SummedStats().Armor)
	// 	})
	// }

	statTable.SetBorder(true)
	statTable.AppendBulk(statData)
	statTable.Render()
}

func (p Player) renderBuffsCard(t io.Writer) {
	buffTable := tablewriter.NewWriter(t)
	buffTable.SetHeader([]string{"Buff", "Enabled"})
	buffData := [][]string{
		{"Blessing Of Kings", fmt.Sprintf("%t", p.PlayerBuffs.BlessingOfKings.Active)},
		{"Blessing Of Might", fmt.Sprintf("%t", p.PlayerBuffs.BlessingOfMight.Active)},
		{"Blessing Of Wisdom", fmt.Sprintf("%t", p.PlayerBuffs.BlessingOfWisdom.Active)},
		{"", ""},
		{"Battle Shout", fmt.Sprintf("%t", p.PlayerBuffs.BattleShout.Active)},
		{"", ""},
		{"Trueshot", fmt.Sprintf("%t", p.PlayerBuffs.TrueShot.Active)},
		{"", ""},
		{"Grace Of AirTotem", fmt.Sprintf("%t", p.PlayerBuffs.GraceOfAirTotem.Active)},
		{"Strength OfEarth Totem", fmt.Sprintf("%t", p.PlayerBuffs.StrengthOfEarthTotem.Active)},
		{"Mana Spring Totem", fmt.Sprintf("%t", p.PlayerBuffs.ManaSpringTotem.Active)},
		{"Windfury Totem (Not Implemented)", fmt.Sprintf("%t", p.PlayerBuffs.WindfuryTotem.Active)},
		{"", ""},
		{"Arcane Brilliance", fmt.Sprintf("%t", p.PlayerBuffs.ArcaneBrilliance.Active)},
		{"", ""},
		{"Gift Of TheWild", fmt.Sprintf("%t", p.PlayerBuffs.GiftOfTheWild.Active)},
		{"Leader Of ThePack", fmt.Sprintf("%t", p.PlayerBuffs.LeaderOfThePack.Active)},
		{"", ""},
		{"Bloodlust", fmt.Sprintf("%d", p.PlayerBuffs.Bloodlust)},
		{"", ""},
		{"Prayer Of Fortitude", fmt.Sprintf("%t", p.PlayerBuffs.PrayerOfFortitude.Active)},
		{"", ""},
		{"Blood Pact", fmt.Sprintf("%t", p.PlayerBuffs.BloodPact.Active)},
		{"", ""},
		{"Braided Eternium Chain", fmt.Sprintf("%t", p.PlayerBuffs.BraidedEterniumChain.Active)},
		{"Ferocious Inspiration", fmt.Sprintf("%t, n=%d, uptime=%f", p.PlayerBuffs.FerociousInspiration.Value > 0, p.PlayerBuffs.FerociousInspiration.Value, p.PlayerBuffs.FerociousInspiration.Uptime)},
	}

	for _, row := range buffData {
		if strings.Contains(row[1], "true") || (row[0] == "Bloodlust" && p.PlayerBuffs.Bloodlust > 0) {
			buffTable.Rich(row, []tablewriter.Colors{nil, {tablewriter.Normal, tablewriter.FgGreenColor}})
		} else {
			buffTable.Rich(row, []tablewriter.Colors{nil, {tablewriter.Normal, tablewriter.FgRedColor}})
		}
	}

	buffTable.Render()
}

func (p Player) renderConsumsCard(t io.Writer) {
	table := tablewriter.NewWriter(t)
	table.SetHeader([]string{"Consumable", "Enabled"})

	data := [][]string{
		{"Leatherworking Drums", fmt.Sprintf("%t", p.ActivatedConsumables.LeatherworkingDrums)},
		{"Haste Potion", fmt.Sprintf("%t", p.ActivatedConsumables.HastePotion)},
		{"Mana Potion", fmt.Sprintf("%t", p.ActivatedConsumables.ManaPotion)},
	}

	for _, row := range data {
		if row[1] == "true" {
			table.Rich(row, []tablewriter.Colors{nil, {tablewriter.Normal, tablewriter.FgGreenColor}})
		} else {
			table.Rich(row, []tablewriter.Colors{nil, {tablewriter.Normal, tablewriter.FgRedColor}})
		}
	}
	table.Render()
}

func (p Player) renderDebuffsCard(t io.Writer) {
	table := tablewriter.NewWriter(t)
	table.SetHeader([]string{"Debuff", "Bonus", "Enabled"})

	data := [][]string{
		{"JoW", "", fmt.Sprintf("%t", p.TargetDebuffs.JudgementOfWisdom.Active)},
		{"JoC", "", fmt.Sprintf("%t", p.TargetDebuffs.JudgementOfTheCrusader.Active)},
		{"Curse of Elements", fmt.Sprintf("%t", p.TargetDebuffs.CurseOfElements.Improved), fmt.Sprintf("%t", p.TargetDebuffs.CurseOfElements.Active)},
		{"Expose Weakness", fmt.Sprintf("Uptime=%f, SVHunterAgi=%d", p.TargetDebuffs.ExposeWeakness.Uptime, p.TargetDebuffs.ExposeWeakness.Value), fmt.Sprintf("%t", p.TargetDebuffs.ExposeWeakness.Active)},
		{"Sunder Armor", "", fmt.Sprintf("%t", p.TargetDebuffs.SunderArmor.Active)},
		{"Expose Armor", fmt.Sprintf("%t", p.TargetDebuffs.ExposeArmor.Improved), fmt.Sprintf("%t", p.TargetDebuffs.ExposeArmor.Active)},
		{"Curse of Recklessness", "", fmt.Sprintf("%t", p.TargetDebuffs.CurseOfRecklessness.Active)},
		{"Faerie Fire", fmt.Sprintf("%t", p.TargetDebuffs.FaeriFire.Improved), fmt.Sprintf("%t", p.TargetDebuffs.FaeriFire.Active)},
		{"Misery", "", fmt.Sprintf("%t", p.TargetDebuffs.Misery.Active)},
		{"Blood Frenzy", "", fmt.Sprintf("%t", p.TargetDebuffs.BloodFrenzy.Active)},
	}

	for _, row := range data {
		if row[2] == "true" {
			table.Rich(row, []tablewriter.Colors{nil, nil, {tablewriter.Normal, tablewriter.FgGreenColor}})
		} else {
			table.Rich(row, []tablewriter.Colors{nil, nil, {tablewriter.Normal, tablewriter.FgRedColor}})
		}
	}

	table.Render()
}

func (p Player) PrintDescription(t io.Writer) {
	p.renderStatsCard(t)
	p.renderBuffsCard(t)
	p.renderConsumsCard(t)
	p.renderDebuffsCard(t)
}
