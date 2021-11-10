package equipment

import (
	"testing"
)

func TestEquipment_TotalArmorPen(t *testing.T) {
	tests := []struct {
		name      string
		equipment Equipment
		want      int
	}{
		{
			name: "basic",
			equipment: Equipment{
				Helm: ArmorItem{
					BaseStats: BaseStats{
						ArmorPenetration: 22,
					},
				},
				Neck:      ArmorItem{},
				Shoulders: ArmorItem{},
				Cloak:     ArmorItem{},
				Chest: ArmorItem{
					BaseStats: BaseStats{
						ArmorPenetration: 37,
					},
				},
				Bracers:    ArmorItem{},
				Gloves:     ArmorItem{},
				Belt:       ArmorItem{},
				Pants:      ArmorItem{},
				Boots:      ArmorItem{},
				RingOne:    ArmorItem{},
				RingTwo:    ArmorItem{},
				Ranged:     RangedWeapon{},
				Quiver:     Quiver{},
				TrinketOne: ArmorItem{},
				TrinketTwo: ArmorItem{},
			},
			want: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.equipment.ArmorPenetration(); got != tt.want {
				t.Errorf("Equipment.TotalArmorPen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquipment_Agility(t *testing.T) {
	tests := []struct {
		name      string
		equipment Equipment
		want      int
	}{
		{
			name: "basic",
			equipment: Equipment{
				Helm: ArmorItem{
					BaseStats: BaseStats{
						Agility: 13,
					},
					Enchant: Enchant{
						Agility: 15,
					},
					Gems: GemSlots{
						SlottedGems: []Gem{
							{
								BaseStats: BaseStats{
									Agility: 8,
								},
							},
						},
					},
				},
				Chest: ArmorItem{
					BaseStats: BaseStats{
						Agility:  55,
						Strength: 12,
					},
				},
				Ranged: RangedWeapon{
					Weapon: Weapon{
						DamageMin: 223,
						DamageMax: 423,
						Speed:     1.7,
						ArmorItem: ArmorItem{
							BaseStats: BaseStats{
								Agility: 61,
								Spirit:  7,
							},
						},
					},
				},
			},
			want: 129 + 8 + 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.equipment.Agility(); got != tt.want {
				t.Errorf("Equipment.Agility() = %v, want %v", got, tt.want)
			}
		})
	}
}
