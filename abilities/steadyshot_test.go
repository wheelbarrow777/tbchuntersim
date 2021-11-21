package abilities

import (
	"tbchuntersim/equipment"
	"tbchuntersim/player"
	"tbchuntersim/util"
	"testing"
)

func TestSteadyShot_Weight(t *testing.T) {
	type args struct {
		p *player.Player
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "no_modifiers",
			args: args{
				p: &player.Player{
					Am: player.ActiveModifiers{},
					Talents: player.Talents{
						BM: player.BM{
							SerpentsSwiftness: 5,
						},
					},
					Equipment: equipment.Equipment{
						Ranged: equipment.RangedWeapon{
							Weapon: equipment.Weapon{
								DamageMin: 217,
								DamageMax: 327,
								Speed:     3,
							},
							AmmoDPS: 43,
						},
						Quiver: equipment.Quiver{
							Speed: 1.15,
						},
					},
				},
			},
			want: 1116.47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := NewSteadyShot()
			got := ss.Weight(tt.args.p)
			if !util.CompFloat(got, tt.want, 200) {
				t.Errorf("SteadyShot.Weight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteadyShot_calcAvgDamage(t *testing.T) {
	type fields struct {
		BaseCastTime    float64
		CurrentCooldown float64
		AverageDamage   float64
	}
	type args struct {
		p *player.Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "no_modifiers",
			fields: fields{
				BaseCastTime:    0.5,
				CurrentCooldown: 0,
				AverageDamage:   0,
			},
			args: args{
				p: &player.Player{
					Am: player.ActiveModifiers{},
					Talents: player.Talents{
						BM: player.BM{
							SerpentsSwiftness: 5,
						},
					},
					Equipment: equipment.Equipment{
						Ranged: equipment.RangedWeapon{
							Weapon: equipment.Weapon{
								DamageMin: 217,
								DamageMax: 327,
								Speed:     3,
							},
							AmmoDPS: 43,
						},
						Quiver: equipment.Quiver{
							Speed: 1.15,
						},
					},
				},
			},
			want: 1320.53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := SteadyShot{
				BaseCastTime:    tt.fields.BaseCastTime,
				CurrentCooldown: tt.fields.CurrentCooldown,
			}
			got := ss.calcAvgDamage(tt.args.p)
			t.Log(got)
			if !util.CompFloat(got, tt.want, 20) {
				t.Errorf("AutoShot.calcAvgDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
