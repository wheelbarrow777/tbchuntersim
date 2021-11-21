package cmd

import (
	"fmt"
	"io"
	"os"
	"tbchuntersim/preset"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	enableAllBuffs  = false
	disableAllBuffs = false

	enableAllDebuffs  = false
	disableAllDebuffs = false

	enableAllConsumables  = false
	disableAllConsumables = false

	modifierProfileFile = ""
)

type modificationWrapper struct {
	Mod  preset.SimulationPreset
	Opts preset.ModificationOptions
}

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:     "modify target_preset.json",
	Short:   "Modify an existing preset",
	Example: "tbchunersim preset modify mypreset.json --enable-all-buffs",
	Long:    ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("first argument must be the target preset file")
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if enableAllBuffs || disableAllBuffs {
			if enableAllBuffs == disableAllBuffs {
				return fmt.Errorf("--all-buffs and --no-buffs can't be set at the same time")
			}
		}

		if enableAllDebuffs || disableAllDebuffs {
			if enableAllDebuffs == disableAllDebuffs {
				return fmt.Errorf("--all-debuffs and --no-debuffs can't be set at the same time")
			}
		}

		if enableAllConsumables || disableAllConsumables {
			if enableAllConsumables == disableAllConsumables {
				return fmt.Errorf("--all-consumables and --no-consumables can't be set at the same time")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		presetLocation := args[0]

		if _, err := os.Stat(presetLocation); err == nil {
			// File Exists, open it and read the content
			f, err := os.Open(presetLocation)
			if err != nil {
				log.WithError(err).Fatal("Could not open the target preset file for readoing")
			}

			existingPreset := preset.SimulationPreset{}
			_, err = io.Copy(&existingPreset, f)
			f.Close()
			if err != nil {
				log.WithError(err).Fatal("Could not parse the existing preset")
			}

			// Build modification queue
			modQueue := []modificationWrapper{}
			// Apply the modifications
			if modifierProfileFile == "" {
				// No buff profile given, modify based on flags
				if enableAllBuffs {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.AllBuffsMod,
						Opts: preset.ModificationOptions{Buffs: true},
					})
				}
				if disableAllBuffs {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.NoBuffsMod,
						Opts: preset.ModificationOptions{Buffs: true},
					})
				}
				if enableAllDebuffs {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.AllDebuffsMod,
						Opts: preset.ModificationOptions{TargetDebuffs: true},
					})
				}
				if disableAllDebuffs {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.NoDebuffsMod,
						Opts: preset.ModificationOptions{TargetDebuffs: true},
					})
				}
				if enableAllConsumables {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.AllConsumablesMod,
						Opts: preset.ModificationOptions{Consumables: true},
					})
				}
				if disableAllConsumables {
					modQueue = append(modQueue, modificationWrapper{
						Mod:  preset.NoConsumablesMod,
						Opts: preset.ModificationOptions{Consumables: true},
					})
				}
			} else {
				// Modifier profile is given, ignore flags and read profile
				f, err = os.Open(modifierProfileFile)
				if err != nil {
					log.WithError(err).Fatal("Cannot open the given modifier profile")
				}
				modProfile := preset.SimulationPreset{}
				_, err := io.Copy(&modProfile, f)
				if err != nil {
					log.WithError(err).Fatal("Cannot read the given modifier profile")
				}
				f.Close()
				modQueue = append(modQueue, modificationWrapper{
					Mod: modProfile,
					Opts: preset.ModificationOptions{
						Consumables:          true,
						ActivatedConsumables: true,
						Buffs:                true,
						TargetDebuffs:        true,
					},
				})
			}

			for _, mod := range modQueue {
				existingPreset.ApplyModification(mod.Mod, mod.Opts)
			}

			// Open the file for writing
			f, err = os.Create(presetLocation)
			if err != nil {
				log.WithError(err).Fatal("Could not open the target preset file for writing")
			}
			defer f.Close()

			_, err = io.Copy(f, existingPreset)
			if err != nil {
				log.WithError(err).Fatal("Could not write the modified preset back")
			}

		} else {
			// File does not exist
			log.WithField("Target File", presetLocation).Fatal("The provided file doesn't exist")
		}

	},
}

func init() {
	presetCmd.AddCommand(modifyCmd)

	modifyCmd.PersistentFlags().BoolVar(&enableAllBuffs, "all-buffs", false, "If set, all buffs in the target preset is enabled")
	modifyCmd.PersistentFlags().BoolVar(&disableAllBuffs, "no-buffs", false, "If set, all buffs in the target preset is disabled")

	modifyCmd.PersistentFlags().BoolVar(&enableAllDebuffs, "all-debuffs", false, "If set, all debuffs in the target preset is disabled")
	modifyCmd.PersistentFlags().BoolVar(&disableAllDebuffs, "no-debuffs", false, "If set, all debuffs in the target preset is disabled")

	modifyCmd.PersistentFlags().BoolVar(&enableAllConsumables, "all-consumables", false, "If set, all consumables in the target preset is disabled")
	modifyCmd.PersistentFlags().BoolVar(&disableAllConsumables, "no-consumables", false, "If set, all consumables in the target preset is disabled")

	modifyCmd.PersistentFlags().StringVarP(&modifierProfileFile, "modifier-profile", "p", "", "If set, ignores all buffs, debuffs and consumables are applied from the given modifier profile. All other flags are ignored. Must point to a preset JSON file.")

}
