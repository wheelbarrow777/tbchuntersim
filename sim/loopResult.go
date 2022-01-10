package sim

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type AbilityDetails struct {
	TotalDamage float64
	NumHits     int
}

type LoopResult struct {
	Time              []float64
	Mana              []float64
	Damage            []float64
	RangedAttackSpeed []float64
	Ability           map[string]AbilityDetails
	MadnessUptimeData []bool
}

func (sr LoopResult) flooredTime() []int {
	rS := []int{}
	for _, v := range sr.Time {
		rS = append(rS, int(v))
	}
	return rS
}

func (sr LoopResult) ManaChart(filename string) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Mana Usage",
	}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < len(sr.Time); i++ {
		lineItems = append(lineItems, opts.LineData{Value: sr.Mana[i]})
	}

	line.SetXAxis(sr.flooredTime()).AddSeries("mana", lineItems)
	f, _ := os.Create(filename)
	line.Render(f)
}

func (sr LoopResult) MadnessUptimeChart(filename string) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Madness Uptime",
	}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < len(sr.Time); i++ {
		if sr.MadnessUptimeData[i] {
			lineItems = append(lineItems, opts.LineData{Value: 1})
		} else {
			lineItems = append(lineItems, opts.LineData{Value: 0})
		}
	}

	line.SetXAxis(sr.flooredTime()).AddSeries("uptime", lineItems)
	f, _ := os.Create(filename)
	line.Render(f)
}

func (sr LoopResult) MadnessUptime() float64 {
	// Divide into second buckets

	var uptimeAtSecond = make([]bool, int(sr.Time[len(sr.Time)-1])+1)
	for i, uptimeTick := range sr.MadnessUptimeData {
		if uptimeTick {
			uptimeAtSecond[int(sr.Time[i])] = true
		}
	}

	// Loop through all buckets, find how many seconds were active
	activeSeconds := 0.0
	for _, s := range uptimeAtSecond {
		if s {
			activeSeconds += 1.0
		}
	}
	return activeSeconds / float64(len(uptimeAtSecond))
}

func (sr LoopResult) DPSAtTimeDeltas(delta int) []float64 {
	accDmg := sr.accumulatedDamage()
	dpsAtDelta := []float64{}
	var deltaDamage float64

	deltaIndex := 1
	timeIndexAtLastDelta := 0
	for i := 0; i < len(sr.Time); i++ {
		if int(sr.Time[i]) >= deltaIndex*delta {
			if timeIndexAtLastDelta == 0 {
				deltaDamage = accDmg[i]
			} else {
				deltaDamage = accDmg[i] - accDmg[timeIndexAtLastDelta-1]
			}
			dpsAtDelta = append(dpsAtDelta, deltaDamage/(sr.Time[i]-float64((deltaIndex-1)*delta)))
			timeIndexAtLastDelta = i
			deltaIndex++
			deltaDamage = 0
		}
	}

	return dpsAtDelta
}

func (sr LoopResult) ManaAtTimeDeltas(delta int) []float64 {
	accMana := sr.accumulatedMana()
	manaAtDelta := []float64{}
	var deltaMana float64

	deltaIndex := 1
	timeIndexAtLastDelta := 0
	numberOfDatapointsInDelta := 1.0
	for i := 0; i < len(sr.Time); i++ {
		if int(sr.Time[i]) >= deltaIndex*delta {
			if timeIndexAtLastDelta == 0 {
				deltaMana = accMana[i]
			} else {
				deltaMana = accMana[i] - accMana[timeIndexAtLastDelta-1]
			}

			manaAtDelta = append(manaAtDelta, deltaMana/numberOfDatapointsInDelta)
			numberOfDatapointsInDelta = 1
			timeIndexAtLastDelta = i
			deltaIndex++
			deltaMana = 0
		}
		numberOfDatapointsInDelta++
	}

	return manaAtDelta
}

func (sr LoopResult) DPSChart(filename string) {
	dpsAtTimeStep := make([]float64, len(sr.Time))
	accDmg := sr.accumulatedDamage()
	for i := 0; i < len(sr.Time); i++ {
		if sr.Time[i] == 0 {
			log.Warn("t=0")
			dpsAtTimeStep[i] = 0
		} else {
			dpsAtTimeStep[i] = accDmg[i] / sr.Time[i]

		}
	}

	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "DPS",
	})).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < len(sr.Time); i++ {
		lineItems = append(lineItems, opts.LineData{Value: dpsAtTimeStep[i]})
	}

	line.SetXAxis(sr.flooredTime()).AddSeries("DPS", lineItems)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	line.Render(f)
}

func (sr LoopResult) RangedASChart(filename string) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Ranged Attack Speed",
	}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < len(sr.Time); i++ {
		lineItems = append(lineItems, opts.LineData{Value: sr.RangedAttackSpeed[i]})
	}

	line.SetXAxis(sr.flooredTime()).AddSeries("eAWS", lineItems)
	f, _ := os.Create(filename)
	line.Render(f)
}

func (sr LoopResult) accumulatedDamage() []float64 {
	accumulatedDamage := make([]float64, len(sr.Time))
	for i := 0; i < len(sr.Time); i++ {
		if i == 0 {
			accumulatedDamage[i] = sr.Damage[i]
		} else {
			accumulatedDamage[i] = sr.Damage[i] + accumulatedDamage[i-1]
		}
	}

	return accumulatedDamage
}

func (sr LoopResult) accumulatedMana() []float64 {
	accumulatedMana := make([]float64, len(sr.Time))
	for i := 0; i < len(sr.Time); i++ {
		if i == 0 {
			accumulatedMana[i] = sr.Mana[i]
		} else {
			accumulatedMana[i] = sr.Mana[i] + accumulatedMana[i-1]
		}
	}

	return accumulatedMana
}

func (sr LoopResult) DPS() float64 {
	totalDmg := 0.0
	for _, d := range sr.Damage {
		totalDmg += d
	}

	return totalDmg / sr.Time[len(sr.Time)-1]
}

func (sr LoopResult) AbilityBreakdownChart(filename string) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Ability Breakdown",
		Subtitle: "Total Damage Dealt",
	}))

	pie.AddSeries("Steady Shot", []opts.PieData{
		{
			Name:  "Steady Shot",
			Value: int(sr.Ability["SteadyShot"].TotalDamage),
		},
		{
			Name:  "Auto Shot",
			Value: int(sr.Ability["AutoShot"].TotalDamage),
		},
	}).SetSeriesOptions(charts.WithLabelOpts(opts.Label{
		Show:      true,
		Formatter: "{b}: {c}",
	}))
	f, err := os.Create(filename)
	if err != nil {
		log.WithError(err).Error("Could not write ability breakdown chart")
		return
	}
	defer f.Close()
	pie.Render(f)
}
