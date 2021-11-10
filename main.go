package main

import (
	"huntsim/config"
	"huntsim/player"
	"huntsim/sim"
	"math/rand"
	"runtime"
	"time"

	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
)

func worker(simConfig *config.SimOptions, jobs <-chan *player.Player, result chan<- float64) {
	for j := range jobs {
		result <- sim.RunSimulationLoop(simConfig, *j)
	}
}

func main() {
	log.Infof<("Starting Slason Hunter Sim with %d workers", runtime.NumCPU())
	rand.Seed(time.Now().UnixMicro())
	// err := config.WriteBaseConfig("myfile.json")
	// if err != nil {
	// 	panic(err)
	// }
	playerConfig, simConfig, err := config.ReadConfig("myfile.json")
	if err != nil {
		panic(err)
	}
	p := player.NewPlayer(playerConfig)

	log.SetLevel(log.InfoLevel)

	const numIts = 100
	bar := progressbar.Default(int64(numIts))
	dps := 0.0
	jobs := make(chan *player.Player, numIts)
	results := make(chan float64, numIts)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(simConfig, jobs, results)
	}

	for j := 1; j <= numIts; j++ {
		jobs <- p
	}
	close(jobs)
	bar.Set(0)

	for a := 1; a <= numIts; a++ {
		dps += <-results
		bar.Add(1)
	}

	// for i := 0; i < numIts; i++ {
	// 	bar.Add(1)
	// 	dps += sim.RunSimulationLoop(simConfig, *p)
	// }
	log.Infof("DPS = %f", dps/float64(numIts))
}
