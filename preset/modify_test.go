package preset

import (
	"reflect"
	"tbchuntersim/player"
	"testing"
)

func TestSimulationPreset_ApplyModification(t *testing.T) {
	type args struct {
		modification SimulationPreset
		opts         ModificationOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all_buffs",
			args: args{
				modification: AllBuffsMod,
				opts: ModificationOptions{
					Buffs: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyPreset := SimulationPreset{}

			emptyPreset.ApplyModification(tt.args.modification, tt.args.opts)

			// Check  that all buffs are applied
			buffs := emptyPreset.Buffs

			value := reflect.ValueOf(buffs)
			for i := 0; i < value.NumField(); i++ {
				field := value.Field(i)
				buffName := value.Type().Field(i).Name
				switch v := field.Interface().(type) {
				case player.Buff:
					if !v.Active {
						t.Errorf("expected %s.Active = true, got %t", buffName, v.Active)
					}
				case player.BuffWithImproved:
					if !v.Active {
						t.Errorf("expected %s.Active = true, got %t", buffName, v.Active)
					}
				case bool:
					if !v {
						t.Errorf("expected %s.Active = true, got %t", buffName, v)
					}
				case int:
					if v <= 0 {
						t.Errorf("expected %s count >= 0, got %d", buffName, v)
					}
				}
			}
		})
	}
}
