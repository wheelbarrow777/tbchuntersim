package cmd

import (
	"io"
	"os"
	"runtime"
	"tbchuntersim/preset"
	"tbchuntersim/sim"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a simulation",
	Long:  `Run a simulation`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read the preset file
		presetFile, err := os.Open(viper.GetString("simulation-preset"))
		if err != nil {
			log.WithError(err).Fatal("Could not open the preset file")
		}
		defer presetFile.Close()

		presetConfig := preset.SimulationPreset{}
		_, err = io.Copy(&presetConfig, presetFile)
		if err != nil {
			log.WithError(err).WithField("preset file", presetFile).Fatal("The given preset file is invalid")
		}

		// Parse the preset
		pConfig, simConfig, err := presetConfig.Parse()
		if err != nil {
			log.WithError(err).Fatal("Could not load the simulation preset")
		}

		if viper.GetInt("chart-bukcket-size") < 5 {
			log.WithField("chart-bukcket-size", viper.GetInt("chart-bukcket-size")).Warn("chart-bukcket-size is too low. This will produce higher DPS charts than reality. Consider increasing the delta to above 5")
		}

		chartsFolder := viper.GetString("charts-folder")
		if chartsFolder[len(chartsFolder)-1] != '/' {
			log.WithField("Charts Folder", chartsFolder).Fatal("Charts folder must have a trailing slash")
		}

		simOpts := sim.RunSimulationOpts{
			PlayerConfig:     pConfig,
			SimulationConfig: simConfig,
			Iterations:       viper.GetInt("iterations"),
			SimWorkers:       viper.GetInt("simulation-workers"),
			ChartTimeDelta:   viper.GetInt("chart-bukcket-size"),
			DisableCharts:    viper.GetBool("disable-charts"),
			ChartsFolder:     viper.GetString("charts-folder"),
			CompDPS:          viper.GetFloat64("compare-dps"),
		}
		log.WithFields(log.Fields{
			"Preset":     viper.GetString("simulation-preset"),
			"Iterations": viper.GetString("iterations"),
			"SimWorkers": viper.GetString("simulation-workers"),
		}).Info("Starting Simulation...")

		sim.RunSimulation(&simOpts)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().String("charts-folder", "charts/", "Folder to save simulation charts in")
	runCmd.PersistentFlags().IntP("iterations", "i", 100, "Number of simulation iterations to be run")
	runCmd.PersistentFlags().IntP("simulation-workers", "w", runtime.NumCPU(), "Number of simulation workers to run in parallel")
	runCmd.PersistentFlags().Int("average-damage-iterations", 50, "Number of iterations used when calculating average damage")
	runCmd.PersistentFlags().Int("chart-bukcket-size", 10, "The time delta of the average plots")
	runCmd.PersistentFlags().Bool("disable-charts", false, "If enabled, no charts will be produced")
	runCmd.PersistentFlags().StringP("simulation-preset", "f", "preset.json", "Simulation preset to be ran")
	runCmd.PersistentFlags().Bool("print-resistance", false, "If enabled, print reistance values with base stats")
	runCmd.PersistentFlags().Float64("compare-dps", 0, "If set, compare the simulated DPS with the provided DPS")

	viper.BindPFlag("simulation-preset", runCmd.PersistentFlags().Lookup("simulation-preset"))
	viper.BindPFlag("iterations", runCmd.PersistentFlags().Lookup("iterations"))
	viper.BindPFlag("simulation-workers", runCmd.PersistentFlags().Lookup("simulation-workers"))
	viper.BindPFlag("average-damage-iterations", runCmd.PersistentFlags().Lookup("average-damage-iterations"))
	viper.BindPFlag("chart-bukcket-size", runCmd.PersistentFlags().Lookup("chart-bukcket-size"))
	viper.BindPFlag("charts-folder", runCmd.PersistentFlags().Lookup("charts-folder"))
	viper.BindPFlag("disable-charts", runCmd.PersistentFlags().Lookup("disable-charts"))
	viper.BindPFlag("print-resistance", runCmd.PersistentFlags().Lookup("print-resistance"))
	viper.BindPFlag("compare-dps", runCmd.PersistentFlags().Lookup("compare-dps"))

}
