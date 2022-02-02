package equipment

import "testing"

func TestGemSlots_bonusAchieved(t *testing.T) {
	type fields struct {
		SlotColors []GemColor
		Gems       []Gem
		Bonus      BaseStats
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "two_red_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, RedGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
					{
						Color: RedGem,
					},
				},
				Bonus: BaseStats{},
			},
			want: true,
		},
		{
			name: "two_red_not_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, RedGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
				},
				Bonus: BaseStats{},
			},
			want: false,
		},
		{
			name: "triple_single_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, YellowGem, BlueGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
					{
						Color: YellowGem,
					},
					{
						Color: BlueGem,
					},
				},
				Bonus: BaseStats{},
			},
			want: true,
		},
		{
			name: "multiple_one_single_rest_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, RedGem, YellowGem, BlueGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
					{
						Color: RedGem,
					},
					{
						Color: YellowGem,
					},
					{
						Color: BlueGem,
					},
				},
				Bonus: BaseStats{},
			},
			want: true,
		},
		{
			name: "mixed_gems_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, YellowGem, BlueGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
					{
						Color: Orange,
					},
					{
						Color: Purple,
					},
				},
				Bonus: BaseStats{},
			},
			want: true,
		},
		{
			name: "mixed_gems_not_achieved",
			fields: fields{
				SlotColors: []GemColor{
					RedGem, YellowGem, BlueGem,
				},
				Gems: []Gem{
					{
						Color: RedGem,
					},
					{
						Color: Purple,
					},
					{
						Color: Purple,
					},
				},
				Bonus: BaseStats{},
			},
			want: false,
		},
		{
			name: "edgewalker",
			fields: fields{
				SlotColors: []GemColor{
					RedGem,
					YellowGem,
				},
				Gems: []Gem{
					{
						Color: YellowGem,
						BaseStats: BaseStats{
							CritRating: 12,
						},
					},
					{
						Color: Purple,
						BaseStats: BaseStats{
							Agility: 4,
							Stamina: 6,
						},
					},
				},
				Bonus: BaseStats{
					HitRating: 3,
				},
			},
			want: true,
		},
		{
			name: "prismatic_gems",
			fields: fields{
				SlotColors: []GemColor{
					RedGem,
					YellowGem,
				},
				Gems: []Gem{
					{
						Color: Prismatic,
						BaseStats: BaseStats{
							Resistance: Resistance{
								Fire: 4,
							},
						},
					},
					{
						Color: Prismatic,
						BaseStats: BaseStats{
							Resistance: Resistance{
								Fire: 4,
							},
						},
					},
				},
				Bonus: BaseStats{
					Resistance: Resistance{
						Arcane: 5,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := GemSlots{
				SlotColors:  tt.fields.SlotColors,
				SlottedGems: tt.fields.Gems,
				Bonus:       tt.fields.Bonus,
			}
			got := gs.bonusAchieved()
			if got != tt.want {
				t.Errorf("GemSlots.bonusAchieved() = %v, want %v", got, tt.want)
			}
		})
	}
}
