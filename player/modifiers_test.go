package player

import (
	"huntsim/util"
	"testing"
)

func TestActiveModifiers_ReduceModifierTime(t *testing.T) {
	type fields struct {
		TimerModifiers timerModifier
	}
	type args struct {
		duration float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic",
			fields: fields{
				TimerModifiers: timerModifier{
					RapidFire:  66.2,
					TBW:        87.1,
					QuickShots: 3.0,
					Bloodlust:  0.0,
				},
			},
			args: args{
				duration: 22.3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := &ActiveModifiers{
				TimerModifiers: tt.fields.TimerModifiers,
			}

			rapidFireExpect := 43.9
			tbwExpect := 64.8
			quickShotsExpect := 0.0
			bloodlustExpected := 0.0

			am.ReduceModifierTime(tt.args.duration)

			if !util.CompFloat(am.TimerModifiers.RapidFire, rapidFireExpect, 0.1) {
				t.Errorf("expected RapidFire = %f, got %f", rapidFireExpect, am.TimerModifiers.RapidFire)
			}
			if !util.CompFloat(am.TimerModifiers.TBW, tbwExpect, 0.1) {
				t.Errorf("expected TBW = %f, got %f", tbwExpect, am.TimerModifiers.RapidFire)
			}
			if !util.CompFloat(am.TimerModifiers.QuickShots, quickShotsExpect, 0.1) {
				t.Errorf("expected QuickShots = %f, got %f", quickShotsExpect, am.TimerModifiers.RapidFire)
			}
			if !util.CompFloat(am.TimerModifiers.Bloodlust, bloodlustExpected, 0.1) {
				t.Errorf("expected Bloodlust = %f, got %f", bloodlustExpected, am.TimerModifiers.RapidFire)
			}
		})
	}
}
