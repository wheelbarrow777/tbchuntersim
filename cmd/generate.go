package cmd

import (
	"bufio"
	"os"
	"strings"
	"tbchuntersim/preset"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate target_preset.json",
	Short: "Generates a new preset",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify the arguments
		if len(args) != 1 {
			log.WithField("# Arguments", len(args)).Fatal("Only one argument can be provided. The only argument should be the target of the preset")
		}

		presetLocation := args[0]
		performWrite := false

		// If the file exists
		if _, err := os.Stat(presetLocation); err == nil {
			log.Infof("The file %s already exists, do you want to overwrite? (y/n)", presetLocation)
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("y", text) == 0 {
				performWrite = true
			} else {
				log.Fatal("Not overwriting, exiting")
			}
		} else {
			performWrite = true
		}

		if performWrite {
			log.Infof("Writing preset to %s", presetLocation)
			// Create the file
			presetFile, err := os.Create(presetLocation)
			if err != nil {
				log.WithError(err).Fatal("Could not create the new preset file")
			}
			defer presetFile.Close()
			preset.Generate(presetFile)
		}
	},
}

func init() {
	presetCmd.AddCommand(generateCmd)
}
