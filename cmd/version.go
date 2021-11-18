package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the current version of the simulator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
