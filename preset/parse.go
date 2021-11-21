package preset

import (
	"fmt"
	"reflect"
	"strings"
	"tbchuntersim/consumables"
	"tbchuntersim/equipment"
	"tbchuntersim/itemdb"
	"tbchuntersim/player"
)

func (parsedConfig SimulationPreset) Parse() (pConfig *player.PlayerConfig, simOpts SimOptions, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			retErr = fmt.Errorf("could not read simulation preset: %s", r)
		}
	}()

	pConfig = &player.PlayerConfig{
		Race:                 strings.ToLower(parsedConfig.Race),
		Talents:              parsedConfig.Talents,
		PlayerBuffs:          player.PlayerBuffs{},
		TargetDebuffs:        parsedConfig.TargetDebuffs,
		Equipment:            equipment.Equipment{},
		StaticConsumables:    consumables.StaticConsumables{},
		ActivatedConsumables: parsedConfig.ActivatedConsumables,
	}

	// Apply Buffs
	buffs := parsedConfig.Buffs
	if buffs.ArcaneBrilliance {
		pConfig.PlayerBuffs.ArcaneBrilliance.Active = true
	}
	pConfig.PlayerBuffs.BattleShout = buffs.BattleShout
	pConfig.PlayerBuffs.BlessingOfKings = buffs.BlessingOfKings
	pConfig.PlayerBuffs.BlessingOfMight = buffs.BlessingOfMight
	pConfig.PlayerBuffs.BlessingOfWisdom = buffs.BlessingOfWisdom
	pConfig.PlayerBuffs.BloodPact = buffs.BloodPact

	pConfig.PlayerBuffs.FerociousInspiration.Uptime = buffs.FerociousInspirationCount.Uptime
	pConfig.PlayerBuffs.FerociousInspiration.Value = buffs.FerociousInspirationCount.ExtraHunters

	pConfig.PlayerBuffs.Bloodlust = buffs.BloodLustCount
	if buffs.BraidedEterniumChain {
		pConfig.PlayerBuffs.BraidedEterniumChain.Active = true
	}
	pConfig.PlayerBuffs.GiftOfTheWild = buffs.GiftOfTheWild
	pConfig.PlayerBuffs.GraceOfAirTotem = buffs.GraceOfAirTotem
	pConfig.PlayerBuffs.LeaderOfThePack = buffs.LeaderOfThePack
	pConfig.PlayerBuffs.ManaSpringTotem = buffs.ManaSpringTotem
	pConfig.PlayerBuffs.PrayerOfFortitude = buffs.PrayerOfFortitude
	pConfig.PlayerBuffs.StrengthOfEarthTotem = buffs.StrengthOfEarthTotem
	if buffs.TrueShotAura {
		pConfig.PlayerBuffs.TrueShot.Active = true
	}
	pConfig.PlayerBuffs.WindfuryTotem.Active = false

	// Apply Consumes
	if parsedConfig.Consumables.Food != "" {
		pConfig.StaticConsumables.Player = append(pConfig.StaticConsumables.Player, itemdb.GetConsumable(parsedConfig.Consumables.Food))
	}
	if parsedConfig.Consumables.BattleElixir != "" {
		pConfig.StaticConsumables.Player = append(pConfig.StaticConsumables.Player, itemdb.GetConsumable(parsedConfig.Consumables.BattleElixir))
	}
	if parsedConfig.Consumables.GuardianElixir != "" {
		pConfig.StaticConsumables.Player = append(pConfig.StaticConsumables.Player, itemdb.GetConsumable(parsedConfig.Consumables.GuardianElixir))
	}
	if parsedConfig.Consumables.AgilityScroll != "" {
		pConfig.StaticConsumables.Player = append(pConfig.StaticConsumables.Player, itemdb.GetConsumable(parsedConfig.Consumables.AgilityScroll))
	}
	if parsedConfig.Consumables.StrengthScroll != "" {
		pConfig.StaticConsumables.Player = append(pConfig.StaticConsumables.Player, itemdb.GetConsumable(parsedConfig.Consumables.StrengthScroll))
	}

	// Add gear
	pConfig.Equipment.Helm = itemdb.GetHelmet(parsedConfig.Equipment.Helm.Name)
	pConfig.Equipment.Neck = itemdb.GetNecklace(parsedConfig.Equipment.Neck.Name)
	pConfig.Equipment.Shoulders = itemdb.GetShoulder(parsedConfig.Equipment.Shoulders.Name)
	pConfig.Equipment.Cloak = itemdb.GetCloak(parsedConfig.Equipment.Cloak.Name)
	pConfig.Equipment.Chest = itemdb.GetChest(parsedConfig.Equipment.Chest.Name)
	pConfig.Equipment.Bracers = itemdb.GetBracers(parsedConfig.Equipment.Bracers.Name)
	pConfig.Equipment.Gloves = itemdb.GetGloves(parsedConfig.Equipment.Gloves.Name)
	pConfig.Equipment.Belt = itemdb.GetBelt(parsedConfig.Equipment.Belt.Name)
	pConfig.Equipment.Pants = itemdb.GetPants(parsedConfig.Equipment.Pants.Name)
	pConfig.Equipment.Boots = itemdb.GetBoots(parsedConfig.Equipment.Boots.Name)
	pConfig.Equipment.RingOne = itemdb.GetRing(parsedConfig.Equipment.RingOne.Name)
	pConfig.Equipment.RingTwo = itemdb.GetRing(parsedConfig.Equipment.RingTwo.Name)
	pConfig.Equipment.MainHand = itemdb.GetMeleeWeapon(parsedConfig.Equipment.MainHand.Name)
	if !pConfig.Equipment.MainHand.IsTwoHanded {
		pConfig.Equipment.OffHand = itemdb.GetMeleeWeapon(parsedConfig.Equipment.OffHand.Name)
	}

	pConfig.Equipment.Ranged = itemdb.GetRangedWeapon(parsedConfig.Equipment.Ranged.Name)
	pConfig.Equipment.TrinketOne = itemdb.GetTrinket(parsedConfig.Equipment.TrinketOne.Name)
	pConfig.Equipment.TrinketTwo = itemdb.GetTrinket(parsedConfig.Equipment.TrinketTwo.Name)
	pConfig.Equipment.Ranged.AmmoDPS = parsedConfig.Equipment.AmmoDPS
	pConfig.Equipment.Quiver = itemdb.GetQuiver(parsedConfig.Equipment.Quiver)

	// Apply Oils
	if parsedConfig.Equipment.MainHand.Oil != "" {
		pConfig.Equipment.MainHand.Oil = itemdb.GetOil(parsedConfig.Equipment.MainHand.Oil)
	}
	if parsedConfig.Equipment.OffHand.Oil != "" {
		pConfig.Equipment.OffHand.Oil = itemdb.GetOil(parsedConfig.Equipment.OffHand.Oil)
	}

	// Apply Enchants and Gems
	parsedValue := reflect.ValueOf(parsedConfig.Equipment)
	pConfigValue := reflect.ValueOf(&pConfig.Equipment)
	for i := 0; i < parsedValue.NumField(); i++ {
		itemField := parsedValue.Field(i)
		if itemField.IsValid() {
			if itemField.Kind() == reflect.Struct {
				itemSlotName := parsedValue.Type().Field(i).Name

				// Apply Enchants
				enchantField := itemField.FieldByName("Enchant")
				if enchantField.IsValid() {
					enchantName := enchantField.Interface().(string)
					enchantSlot := pConfigValue.Elem().FieldByName(itemSlotName).FieldByName("Enchant")
					if enchantName != "" {
						enchantSlot.Set(reflect.ValueOf(itemdb.GetEnchant(enchantName)))
					}
				}

				// Apply gems
				gemsField := itemField.FieldByName("Gems")
				if gemsField.IsValid() {
					gems := gemsField.Interface().([]string)
					gemSlotsValue := pConfigValue.Elem().FieldByName(itemSlotName).FieldByName("Gems").Addr()
					if gemSlotsValue.IsValid() {
						addGemMethodValue := gemSlotsValue.MethodByName("AddGem")

						for _, gemName := range gems {
							in := []reflect.Value{reflect.ValueOf(itemdb.GetGem(gemName))}
							addGemMethodValue.Call(in)
						}
					}
				}
			}

		}
	}

	// Apply Scope
	if parsedConfig.Equipment.Ranged.Scope != "" {
		pConfig.Equipment.Ranged.Scope = itemdb.GetScope(parsedConfig.Equipment.Ranged.Scope)
	}

	return pConfig, parsedConfig.Options, retErr
}
