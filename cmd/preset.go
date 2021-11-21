package cmd

import (
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:   "preset",
	Short: "Create, modify and read presets",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(presetCmd)
}
