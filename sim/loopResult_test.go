package sim

import (
	"reflect"
	"testing"
)

func TestSimulationLoopResult_accumulatedDamage(t *testing.T) {
	type fields struct {
		Time    []float64
		Mana    []float64
		Damage  []float64
		Ability map[string]AbilityDetails
	}
	tests := []struct {
		name   string
		fields fields
		want   []float64
	}{
		{
			name: "small_increase",
			fields: fields{
				Time: []float64{
					0.34, 1.23, 3.33, 6.66, 8.12,
				},
				Damage: []float64{
					1342, 1002, 1111, 677, 1892,
				},
			},
			want: []float64{
				1342,
				1342 + 1002,
				1342 + 1002 + 1111,
				1342 + 1002 + 1111 + 677,
				1342 + 1002 + 1111 + 677 + 1892,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := LoopResult{
				Time:    tt.fields.Time,
				Mana:    tt.fields.Mana,
				Damage:  tt.fields.Damage,
				Ability: tt.fields.Ability,
			}
			if got := sr.accumulatedDamage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationLoopResult.accumulatedDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
