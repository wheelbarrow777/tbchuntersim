package preset

import (
	"encoding/json"
	"fmt"
	"io"
)

func (preset *SimulationPreset) Write(p []byte) (n int, err error) {
	err = json.Unmarshal(p, preset)
	if err != nil {
		return 0, fmt.Errorf("given data does not fit into a preset")
	}

	return len(p), nil
}

func (preset SimulationPreset) Read(p []byte) (n int, err error) {
	d, err := json.MarshalIndent(preset, "", "	")
	if err != nil {
		return 0, fmt.Errorf("could not read simulation preset: %w", err)
	}

	n = copy(p, d)
	return n, io.EOF
}

func Generate(writer io.Writer) error {
	base := SimulationPreset{}

	// Apply defaults
	base.Buffs.FerociousInspirationCount.Uptime = 0.94

	_, err := io.Copy(writer, base)

	return err
}
