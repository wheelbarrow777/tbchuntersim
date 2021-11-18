package sim

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

type ResultConnection struct {
	results []LoopResult
}

func NewResultColleciton() *ResultConnection {
	return &ResultConnection{}
}

func (col *ResultConnection) Add(r LoopResult) {
	col.results = append(col.results, r)
}

func (col ResultConnection) DPS() float64 {
	dps := 0.0
	for _, r := range col.results {
		dps += r.DPS()
	}

	return dps / float64(len(col.results))
}

func inlineSum(x *[]float64, y []float64) {
	for i, v := range y {
		(*x)[i] += v
	}
}

func inlineDivide(x *[]float64, denominator float64) {
	for i, v := range *x {
		(*x)[i] = v / denominator
	}
}

func (col ResultConnection) generateDataDelta(delta int, dataType string, stopTime int) ([]float64, error) {
	// Allocate the result array
	summedDeltaValue := make([]float64, int(stopTime)/delta)

	// For each result, get the time delta
	var deltaValue []float64
	for _, r := range col.results {
		switch dataType {
		case "dps":
			deltaValue = r.DPSAtTimeDeltas(delta)
		case "mana":
			deltaValue = r.ManaAtTimeDeltas(delta)
		default:
			return nil, fmt.Errorf("the given dataType is not valid")
		}
		if len(summedDeltaValue) != len(deltaValue) {
			log.WithFields(log.Fields{
				"len(deltaDPS)":       len(deltaValue),
				"len(summedDeltaDPS)": len(summedDeltaValue),
			}).Fatal("Index mismatch when summing up data for DPS Chart")
		}

		inlineSum(&summedDeltaValue, deltaValue)
	}

	// Divide all the DPS entries with the number of iterations to get the average
	inlineDivide(&summedDeltaValue, float64(len(col.results)))

	return summedDeltaValue, nil
}

func (col ResultConnection) deltaTimes(delta int) (float64, []float64) {
	// Allocate the result array
	stopTime := col.results[0].Time[len(col.results[0].Time)-1]
	deltaTimes := make([]float64, int(stopTime)/delta)

	// Create the deltaTimes
	for i := 0; i < len(deltaTimes); i++ {
		deltaTimes[i] = float64(i * delta)
	}

	return stopTime, deltaTimes
}

func (col ResultConnection) ManaChart(delta int, filename string) {
	stopTime, deltaTimes := col.deltaTimes(delta)

	summedDeltaMana, err := col.generateDataDelta(delta, "mana", int(stopTime))
	if err != nil {
		log.WithError(err).Fatal("The given datatype is not known")
	}

	// Plot the chart
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Average Mana",
			Subtitle: fmt.Sprintf("DeltaT = %d, Iterations = %d", delta, len(col.results)),
		}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < int(stopTime)/delta; i++ {
		lineItems = append(lineItems, opts.LineData{Value: summedDeltaMana[i]})
	}

	// Create the x-axis
	line.SetXAxis(deltaTimes).AddSeries("Mana", lineItems)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	line.Render(f)
}

func (col ResultConnection) DPSChart(delta int, filename string) {
	// Allocate the result array
	stopTime, deltaTimes := col.deltaTimes(delta)

	summedDeltaDPS, err := col.generateDataDelta(delta, "dps", int(stopTime))
	if err != nil {
		log.WithError(err).Fatal("The given datatype is not known")
	}

	// Plot the chart
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Average DPS",
			Subtitle: fmt.Sprintf("DeltaT = %d, Iterations = %d", delta, len(col.results)),
		}))

	lineItems := make([]opts.LineData, 0)
	for i := 0; i < int(stopTime)/delta; i++ {
		lineItems = append(lineItems, opts.LineData{Value: summedDeltaDPS[i]})
	}

	// Create the x-axis
	line.SetXAxis(deltaTimes).AddSeries("DPS", lineItems)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	line.Render(f)
}

func (col ResultConnection) AbilityBreakdownChart(filename string) {
	ssDmg := 0.0
	asDmg := 0.0
	for _, result := range col.results {
		for name, details := range result.Ability {
			if name == "SteadyShot" {
				ssDmg += details.TotalDamage
			} else if name == "AutoShot" {
				asDmg += details.TotalDamage
			}
		}
	}

	ssDmg = ssDmg / float64(len(col.results))
	asDmg = asDmg / float64(len(col.results))

	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Average Ability Breakdown",
		Subtitle: fmt.Sprintf("Total Damage Dealt. Iterations = %d", len(col.results)),
	}))

	pie.AddSeries("Steady Shot", []opts.PieData{
		{
			Name:  "Steady Shot",
			Value: int(ssDmg),
		},
		{
			Name:  "Auto Shot",
			Value: int(asDmg),
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

func (col ResultConnection) AbilityBreakdownTable(writer io.Writer) {
	ssDmg := 0.0
	ssHits := 0
	asDmg := 0.0
	asHits := 0
	kcHits := 0
	for _, result := range col.results {
		for name, details := range result.Ability {
			if name == "SteadyShot" {
				ssDmg += details.TotalDamage
				ssHits += details.NumHits
			} else if name == "AutoShot" {
				asDmg += details.TotalDamage
				asHits += details.NumHits
			} else if name == "Kill Command" {
				kcHits += details.NumHits
			}
		}
	}

	ssDmg = ssDmg / float64(len(col.results))
	asDmg = asDmg / float64(len(col.results))
	ssHits = ssHits / len(col.results)
	asHits = asHits / len(col.results)

	ssAvgHit := int(ssDmg / float64(ssHits))
	asAvgHit := int(asDmg / float64(asHits))

	ssDPS := ssDmg / col.results[0].Time[len(col.results[0].Time)-1]
	asDPS := asDmg / col.results[0].Time[len(col.results[0].Time)-1]

	table := tablewriter.NewWriter(writer)
	data := [][]string{
		{"Steady Shot", fmt.Sprintf("%.1f", ssDmg), fmt.Sprintf("%d", ssHits), fmt.Sprintf("%d", ssAvgHit), fmt.Sprintf("%.1f", ssDPS)},
		{"Auto Shot", fmt.Sprintf("%.1f", asDmg), fmt.Sprintf("%d", asHits), fmt.Sprintf("%d", asAvgHit), fmt.Sprintf("%.1f", asDPS)},
	}

	table.SetHeader([]string{"Ability", "Damage", "Num Hits", "Avg Hit", "DPS"})
	table.SetFooter([]string{"", fmt.Sprintf("%.1f", ssDmg+asDmg), fmt.Sprintf("%d", ssHits+asHits), fmt.Sprintf("%d", ssAvgHit+asAvgHit), fmt.Sprintf("%.1f", ssDPS+asDPS)})
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}
