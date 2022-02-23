package sim

import (
	"fmt"
	"os"
	"tbchuntersim/player"
	"tbchuntersim/preset"

	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
)

func worker(simConfig preset.SimOptions, jobs <-chan *player.Player, result chan<- *LoopResult) {
	for j := range jobs {
		result <- RunSimulationLoop(simConfig, *j)
	}
}

type RunSimulationOpts struct {
	PlayerConfig     *player.PlayerConfig
	SimulationConfig preset.SimOptions
	Iterations       int
	SimWorkers       int
	ChartTimeDelta   int
	DisableCharts    bool
	ChartsFolder     string
	CompDPS          float64
}

func RunSimulation(opts *RunSimulationOpts) {
	p := player.NewPlayer(opts.PlayerConfig)

	p.PrintDescription(os.Stdout)

	numIts := opts.Iterations
	bar := progressbar.Default(int64(numIts))
	jobs := make(chan *player.Player, numIts)
	results := make(chan *LoopResult, numIts)

	for w := 1; w <= opts.SimWorkers; w++ {
		go worker(opts.SimulationConfig, jobs, results)
	}

	for j := 1; j <= numIts; j++ {
		jobs <- p
	}
	close(jobs)
	bar.Set(0)

	col := NewResultColleciton()
	for a := 1; a <= numIts; a++ {
		r := <-results
		if !opts.DisableCharts {
			r.DPSChart(fmt.Sprintf("%ssingle_dps_%d.html", opts.ChartsFolder, a))
			r.ManaChart(fmt.Sprintf("%ssingle_mana_%d.html", opts.ChartsFolder, a))
			r.RangedASChart(fmt.Sprintf("%ssingle_ras_%d.html", opts.ChartsFolder, a))
			r.AbilityBreakdownChart(fmt.Sprintf("%ssingle_ability_breakdown_%d.html", opts.ChartsFolder, a))
			if p.Equipment.TrinketOne.Name == "madness of the betrayer" || p.Equipment.TrinketTwo.Name == "madness of the betrayer" {
				r.MadnessUptimeChart(fmt.Sprintf("%ssingle_madness_uptime_%d.html", opts.ChartsFolder, a))
			}
		}
		col.Add(*r)
		bar.Add(1)
	}

	if !opts.DisableCharts {
		col.DPSChart(opts.ChartTimeDelta, fmt.Sprintf("%sdps_average.html", opts.ChartsFolder))
		col.ManaChart(opts.ChartTimeDelta, fmt.Sprintf("%smana_average.html", opts.ChartsFolder))
		col.AbilityBreakdownChart(fmt.Sprintf("%sability_breakdown.html", opts.ChartsFolder))
	}

	fmt.Printf("\n\n")
	col.AbilityBreakdownTable(os.Stdout)
	fmt.Printf("\n\n")

	if p.Equipment.TrinketOne.Name == "madness of the betrayer" || p.Equipment.TrinketTwo.Name == "madness of the betrayer" {
		log.Infof("Madness Uptime %f", col.MadnessUptime())
	}

	DPS := col.DPS()

	if opts.CompDPS == 0.0 {
		log.Infof("Simulation Complete! DPS = %.2f\n\n", DPS)
	} else {
		diff := ((DPS - opts.CompDPS) / DPS) * 100.0
		log.Infof("Simulation Complete! DPS = %.2f (%.2f%% change)\n\n", DPS, diff)
	}

}
