package abilities

import (
	"huntsim/equipment"
	"huntsim/player"
	"huntsim/util"
	"testing"
)

func TestAutoShot_Weight(t *testing.T) {
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
			want: 3383.33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := AutoShot{
				BaseCastTime:    tt.fields.BaseCastTime,
				CurrentCooldown: tt.fields.CurrentCooldown,
			}
			got := as.Weight(tt.args.p)
			if !util.CompFloat(got, tt.want, 200) {
				t.Errorf("AutoShot.Weight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoShot_calcAvgDamage(t *testing.T) {
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
			want: 1370.04,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := AutoShot{
				BaseCastTime:    tt.fields.BaseCastTime,
				CurrentCooldown: tt.fields.CurrentCooldown,
			}
			got := as.calcAvgDamage(tt.args.p)
			t.Log(got)
			if !util.CompFloat(got, tt.want, 100) {
				t.Errorf("AutoShot.calcAvgDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
