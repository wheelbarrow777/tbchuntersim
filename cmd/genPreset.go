package cmd

import (
	"huntsim/config"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genPresetCmd = &cobra.Command{
	Use:   "genPreset",
	Short: "Generate a simulation preset",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Writing preset to %s", viper.GetString("target"))
		config.WriteBaseConfig(viper.GetString("target"))
	},
}

func init() {
	rootCmd.AddCommand(genPresetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	genPresetCmd.PersistentFlags().String("target", "preset.json", "Target location to write the preset")
	viper.BindPFlag("target", genPresetCmd.PersistentFlags().Lookup("target"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genPresetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
