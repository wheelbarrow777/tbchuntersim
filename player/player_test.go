package player

import (
	"huntsim/consumables"
	"huntsim/equipment"
	"huntsim/itemdb"
	"huntsim/util"
	"testing"
)

func TestPlayer_baseHaste(t *testing.T) {
	type fields struct {
		BaseRangedAP    int
		CritChance      float64
		MissChance      float64
		Equipment       equipment.Equipment
		Talents         Talents
		BaseHasteRating float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "with_swiftness",
			fields: fields{
				BaseHasteRating: 0,
				BaseRangedAP:    1844,
				MissChance:      0.0323,
				Equipment: equipment.Equipment{
					Ranged: equipment.RangedWeapon{
						Weapon: equipment.Weapon{
							DamageMin: 169,
							DamageMax: 314,
							Speed:     2.9,
						},
						AmmoDPS: 43,
					},
					Quiver: equipment.Quiver{
						Speed: 1.15,
					},
				},
				Talents: Talents{
					BM: BM{
						SerpentsSwiftness: 5,
					},
					MM: MM{},
					SV: SV{},
				},
			},
			want: 1.38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Player{
				rangedAttackPower: tt.fields.BaseRangedAP,
				Equipment:         tt.fields.Equipment,
				Talents:           tt.fields.Talents,
				hasteRating:       tt.fields.BaseHasteRating,
			}
			got := p.BaseHaste()
			if !util.CompFloat(got, tt.want, 0.1) {
				t.Errorf("Player.baseHaste() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_CritMissDmg(t *testing.T) {
	tests := []struct {
		name string
		p    Player
		want struct {
			critChance   float64
			missChance   float64
			critModifier float64
		}
	}{
		{
			name: "basic",
			p: Player{
				strength:   207,
				agility:    787,
				stamina:    788,
				intellect:  276,
				spirit:     143,
				critRating: 183,
				hitRating:  138,
				Equipment: equipment.Equipment{
					Helm: equipment.ArmorItem{
						Gems: equipment.GemSlots{
							SlottedGems: []equipment.Gem{
								{
									Name: "relentless earthstorm diamond",
								},
							},
						},
					},
				},
				Talents: Talents{
					MM: MM{
						LethalShots: 5,
						MortalShots: 5,
					},
				},
				PlayerBuffs: PlayerBuffs{
					LeaderOfThePack: Buff{
						Active: true,
					},
				},
			},
			want: struct {
				critChance   float64
				missChance   float64
				critModifier float64
			}{
				critChance:   0.3163,
				missChance:   1 - 0.9976,
				critModifier: 2.378,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.p.CritChance()
			if !util.CompFloat(got, tt.want.critChance, 0.01) {
				t.Errorf("Player.CritChance() = %v, want %v", got, tt.want.critChance)
			}

			got = tt.p.MissChance()
			if !util.CompFloat(got, tt.want.missChance, 0.01) {
				t.Errorf("Player.MissChance() = %v, want %v", got, tt.want.missChance)
			}

			got = tt.p.RangeCritDamageModifier()
			if !util.CompFloat(got, tt.want.critModifier, 0.01) {
				t.Errorf("Player.RangeCritDamageModifier() = %v, want %v", got, tt.want.critModifier)

			}
		})
	}
}

func buildBasicTestConfig() PlayerConfig {
	c := PlayerConfig{
		Race: "orc",
		Talents: Talents{
			BM: BM{
				SerpentsSwiftness: 5,
			},
			MM: MM{
				TrueshotAura: 0,
				LethalShots:  5,
				MortalShots:  5,
			},
		},
		Equipment: equipment.Equipment{
			Helm:      itemdb.GetHelmet("rift stalker helm"),
			Neck:      itemdb.GetNecklace("choker of vile intent"),
			Shoulders: itemdb.GetShoulder("rift stalker mantle"),
			Cloak:     itemdb.GetCloak("cloak of the pit stalker"),
			Chest:     itemdb.GetChest("rift stalker hauberk"),
			Bracers:   itemdb.GetBracers("vambraces of ending"),
			Gloves:    itemdb.GetGloves("rift stalker gauntlets"),
			Belt:      itemdb.GetBelt("belt of deep shadow"),
			Pants:     itemdb.GetPants("void reaver greaves"),
			Boots:     itemdb.GetBoots("edgewalker longboots"),
			RingOne:   itemdb.GetRing("ring of the recalcitrant"),
			RingTwo:   itemdb.GetRing("ring of a thousand marks"),
			Ranged:    itemdb.GetRangedWeapon("sunfury bow of the phoenix"),
			MainHand:  itemdb.GetMeleeWeapon("stellaris"),
			OffHand:   itemdb.GetMeleeWeapon("stormreaver warblades"),
			Quiver: equipment.Quiver{
				Speed: 1.15,
			},
			TrinketOne: itemdb.GetTrinket("bloodlust brooch"),
			TrinketTwo: itemdb.GetTrinket("dragonspine trophy"),
		},
	}

	// Apply enchants
	c.Equipment.Helm.Enchant = itemdb.GetEnchant("glyph of ferocity")
	c.Equipment.Shoulders.Enchant = itemdb.GetEnchant("Greater Inscription of Vengeance")
	c.Equipment.Cloak.Enchant = itemdb.GetEnchant("Enchant Cloak - Greater Agility")
	c.Equipment.Chest.Enchant = itemdb.GetEnchant("Enchant Chest - Exceptional Stats")
	c.Equipment.Bracers.Enchant = itemdb.GetEnchant("Enchant Bracer - Assault")
	c.Equipment.Gloves.Enchant = itemdb.GetEnchant("Enchant Gloves - Superior Agility")
	c.Equipment.Pants.Enchant = itemdb.GetEnchant("Nethercobra Leg Armor")
	c.Equipment.Boots.Enchant = itemdb.GetEnchant("Enchant Boots - Dexterity")
	c.Equipment.Ranged.Scope = itemdb.GetScope("Stabilized Eternium Scope")
	c.Equipment.MainHand.Enchant = itemdb.GetEnchant("Enchant Weapon - Agility")
	c.Equipment.OffHand.Enchant = itemdb.GetEnchant("Enchant Weapon - Agility")

	// Apply Gems
	c.Equipment.Helm.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Helm.Gems.AddGem(itemdb.GetGem("Relentless Earthstorm Diamond"))

	c.Equipment.Shoulders.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Shoulders.Gems.AddGem(itemdb.GetGem("delicate living ruby"))

	c.Equipment.Chest.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Chest.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Chest.Gems.AddGem(itemdb.GetGem("delicate living ruby"))

	c.Equipment.Bracers.Gems.AddGem(itemdb.GetGem("delicate living ruby"))

	c.Equipment.Belt.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Belt.Gems.AddGem(itemdb.GetGem("delicate living ruby"))

	c.Equipment.Pants.Gems.AddGem(itemdb.GetGem("delicate living ruby"))
	c.Equipment.Pants.Gems.AddGem(itemdb.GetGem("Wicked Noble Topaz"))
	c.Equipment.Pants.Gems.AddGem(itemdb.GetGem("Shifting Nightseye"))

	c.Equipment.Boots.Gems.AddGem(itemdb.GetGem("shifting nightseye"))
	c.Equipment.Boots.Gems.AddGem(itemdb.GetGem("stone of blades"))

	return c
}

func TestNewPlayerBase(t *testing.T) {
	c := buildBasicTestConfig()
	p := NewPlayer(&c)

	// Check all the stats
	if p.strength != 73 {
		t.Errorf("expected strength = 73, got %d", p.strength)
	}
	if p.agility != 673 {
		t.Errorf("expected agility = 673, got %d", p.agility)
	}
	if p.stamina != 481 {
		t.Errorf("expected stamina = 481, got %d", p.stamina)
	}
	if p.intellect != 192 {
		t.Errorf("expected intellect = 192, got %d", p.intellect)
	}
	if p.spirit != 92 {
		t.Errorf("expected spirit = 92, got %d", p.spirit)
	}
	if p.rangedAttackPower != 2022 {
		t.Errorf("expected rangedAttackPower = 2022, got %d", p.rangedAttackPower)
	}
	if p.hitRating != 138 {
		t.Errorf("expected hitRating = 138, got %d", p.hitRating)
	}
	if p.critRating != 134 {
		t.Errorf("expected critRating = 134, got %d", p.critRating)
	}
}

func TestNewPlayerBuffsFull(t *testing.T) {
	c := buildBasicTestConfig()
	c.PlayerBuffs = PlayerBuffs{
		BlessingOfKings: Buff{
			Active:   true,
			Improved: true,
		},
		BlessingOfMight: Buff{
			Active:   true,
			Improved: true,
		},
		BlessingOfWisdom: Buff{
			Active:   true,
			Improved: true,
		},
		BattleShout: Buff{
			Active:   true,
			Improved: true,
		},
		TrueShot: Buff{
			Active:   true,
			Improved: true,
		},
		LeaderOfThePack: Buff{
			Active:   true,
			Improved: true,
		},
		GraceOfAirTotem: Buff{
			Active:   false,
			Improved: true,
		},
		StrengthOfEarthTotem: Buff{
			Active:   true,
			Improved: true,
		},
		ManaSpringTotem: Buff{
			Active:   false,
			Improved: true,
		},
		ArcaneBrilliance: Buff{
			Active:   true,
			Improved: true,
		},
		GiftOfTheWild: Buff{
			Active:   true,
			Improved: true,
		},
		PrayerOfFortitude: Buff{
			Active:   true,
			Improved: true,
		},

		BraidedEterniumChain: Buff{
			Active:   true,
			Improved: true,
		},
	}

	p := NewPlayer(&c)

	// Check all the stats
	if p.strength != 207 {
		t.Errorf("expected strength = 207, got %d", p.strength)
	}
	if p.agility != 760 {
		t.Errorf("expected agility = 760, got %d", p.agility)
	}
	if p.stamina != 661 {
		t.Errorf("expected stamina = 661, got %d", p.stamina)
	}
	if p.intellect != 275 {
		t.Errorf("expected intellect = 275, got %d", p.intellect)
	}
	if p.spirit != 121 {
		t.Errorf("expected spirit = 121, got %d", p.spirit)
	}
	if p.meleeAttackPower != 2921 {
		t.Errorf("expected meleeAttackPower = 2921, got %d", p.rangedAttackPower)
	}
	if p.rangedAttackPower != 2498 {
		t.Errorf("expected rangedAttackPower = 2498, got %d", p.rangedAttackPower)
	}
	if p.hitRating != 138 {
		t.Errorf("expected hitRating = 138, got %d", p.hitRating)
	}
	if p.critRating != 162 {
		t.Errorf("expected critRating = 162, got %d", p.critRating)
	}

}

func TestNewPlayerBuffsConsumesFull(t *testing.T) {
	c := buildBasicTestConfig()
	c.PlayerBuffs = PlayerBuffs{
		BlessingOfKings: Buff{
			Active:   true,
			Improved: true,
		},
		BlessingOfMight: Buff{
			Active:   true,
			Improved: true,
		},
		BlessingOfWisdom: Buff{
			Active:   true,
			Improved: true,
		},
		BattleShout: Buff{
			Active:   true,
			Improved: true,
		},
		TrueShot: Buff{
			Active:   true,
			Improved: true,
		},
		LeaderOfThePack: Buff{
			Active:   true,
			Improved: true,
		},
		GraceOfAirTotem: Buff{
			Active:   false,
			Improved: true,
		},
		StrengthOfEarthTotem: Buff{
			Active:   true,
			Improved: true,
		},
		ManaSpringTotem: Buff{
			Active:   false,
			Improved: true,
		},
		ArcaneBrilliance: Buff{
			Active:   true,
			Improved: true,
		},
		GiftOfTheWild: Buff{
			Active:   true,
			Improved: true,
		},
		PrayerOfFortitude: Buff{
			Active:   true,
			Improved: true,
		},

		BraidedEterniumChain: Buff{
			Active:   true,
			Improved: true,
		},
	}
	c.StaticConsumables = consumables.StaticConsumables{
		Player: []consumables.StaticConsumable{
			itemdb.GetConsumable("warp burger"),
			itemdb.GetConsumable("elixir of major agility"),
			itemdb.GetConsumable("elixir of major mageblood"),
			itemdb.GetConsumable("scroll of agility V"),
			itemdb.GetConsumable("scroll of strength V"),
		},
	}

	p := NewPlayer(&c)

	// Check all the stats
	if p.strength != 229 {
		t.Errorf("expected strength = 73, got %d", p.strength)
	}
	if p.agility != 842 {
		t.Errorf("expected agility = 673, got %d", p.agility)
	}
	if p.stamina != 661 {
		t.Errorf("expected stamina = 481, got %d", p.stamina)
	}
	if p.intellect != 275 {
		t.Errorf("expected intellect = 192, got %d", p.intellect)
	}
	if p.spirit != 143 {
		t.Errorf("expected spirit = 92, got %d", p.spirit)
	}
	if p.meleeAttackPower != 3025 {
		t.Errorf("expected meleeAttackPower = 2353, got %d", p.rangedAttackPower)
	}
	if p.rangedAttackPower != 2580 {
		t.Errorf("expected rangedAttackPower = 2353, got %d", p.rangedAttackPower)
	}
	if p.hitRating != 138 {
		t.Errorf("expected hitRating = 138, got %d", p.hitRating)
	}
	if p.critRating != 182 {
		t.Errorf("expected critRating = 134, got %d", p.critRating)
	}
	if !util.CompFloat(p.MaxMana, 7228.0, 0.1) {
		t.Errorf("expected mana = %f, got %f", 7228.0, p.MaxMana)
	}
}
