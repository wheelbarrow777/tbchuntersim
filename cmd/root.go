package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tbchuntersim",
	Short: "A simulator for TBC Hunters",
	Long:  `A simualtor for TBC Hunters`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.huntsim.yaml)")
	rootCmd.PersistentFlags().String("log-level", "INFO", "Desired logging level (INFO, DEBUG, WARNING, ERROR, TRACE)")

	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".huntsim" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".huntsim")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	switch viper.GetString("log-level") {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		fmt.Println("DEBUG")
		log.SetLevel(log.DebugLevel)
	case "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	default:
		log.WithField("submitted log level", viper.GetString("log-level")).Error("Invalid log level provided. Defaulting to DEBUG")
		log.SetLevel(log.DebugLevel)
	}
}
