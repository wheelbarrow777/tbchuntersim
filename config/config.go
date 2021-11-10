package config

import (
	"encoding/json"
	"huntsim/consumables"
	"huntsim/equipment"
	"huntsim/itemdb"
	"huntsim/player"
	"os"
	"reflect"
	"strings"
)

type itemBase struct {
	Name    string
	Gems    []string
	Enchant string
}

type weapon struct {
	itemBase
	Oil string
}

type rangedWeapon struct {
	itemBase
	Scope string
}

type eq struct {
	Helm       itemBase
	Neck       itemBase
	Shoulders  itemBase
	Cloak      itemBase
	Chest      itemBase
	Bracers    itemBase
	Gloves     itemBase
	Belt       itemBase
	Pants      itemBase
	Boots      itemBase
	RingOne    itemBase
	RingTwo    itemBase
	TrinketOne itemBase
	TrinketTwo itemBase
	MainHand   weapon
	OffHand    weapon
	Ranged     rangedWeapon
	Quiver     string
	AmmoDPS    int
}

type consums struct {
	Food           string
	BattleElixir   string
	GuardianElixir string
	AgilityScroll  string
	StrengthScroll string
	UseHastePotion bool
	UseManaPotion  bool

	PetFood           string
	PetScrollAgility  bool
	PetScrollStrenght bool
}

type buff struct {
	Active   bool
	Improved bool
}

type buffs struct {
	BlessingOfKings      player.Buff
	BlessingOfMight      player.Buff
	BlessingOfWisdom     player.Buff
	BattleShout          player.Buff
	TrueShotAura         bool
	LeaderOfThePack      player.Buff
	GraceOfAirTotem      player.Buff
	StrengthOfEarthTotem player.Buff
	ManaSpringTotem      player.Buff
	// WindfuryTotem buff // Not supported
	ArcaneBrilliance     bool
	GiftOfTheWild        player.Buff
	BloodLustCount       int
	LeatherworkingDrums  bool
	PrayerOfFortitude    player.Buff
	BloodPact            player.Buff
	BraidedEterniumChain bool
}

type SimOptions struct {
	SimDuration float64
	Latency     float64
	TargetArmor float64
}

type SimulationConfig struct {
	Race        string
	Equipment   eq
	Consumables consums
	Buffs       buffs
	Talents     player.Talents
	Options     SimOptions
}

func WriteBaseConfig(filename string) error {
	base := SimulationConfig{}
	d, err := json.Marshal(base)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(d)

	return err
}

func ReadConfig(filename string) (*player.PlayerConfig, *SimOptions, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	parsedConfig := SimulationConfig{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&parsedConfig); err != nil {
		return nil, nil, err
	}

	pConfig := player.PlayerConfig{
		Race:              strings.ToLower(parsedConfig.Race),
		Talents:           parsedConfig.Talents,
		PlayerBuffs:       player.PlayerBuffs{},
		TargetDebuffs:     player.TargetDebuffs{},
		Equipment:         equipment.Equipment{},
		StaticConsumables: consumables.StaticConsumables{},
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
	pConfig.PlayerBuffs.Bloodlust = buffs.BloodLustCount
	if buffs.BraidedEterniumChain {
		pConfig.PlayerBuffs.BraidedEterniumChain.Active = true
	}
	if buffs.LeatherworkingDrums {
		pConfig.PlayerBuffs.Drums.Active = true
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

	return &pConfig, &parsedConfig.Options, nil
}
