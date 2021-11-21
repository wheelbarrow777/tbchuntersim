package preset

type ModificationOptions struct {
	Race                 bool
	Equipment            bool
	Consumables          bool
	ActivatedConsumables bool
	Buffs                bool
	TargetDebuffs        bool
	Talents              bool
	Options              bool
}

func (preset *SimulationPreset) ApplyModification(modification SimulationPreset, opts ModificationOptions) *SimulationPreset {
	if opts.Race {
		preset.Race = modification.Race
	}
	if opts.Equipment {
		preset.Equipment = modification.Equipment
	}
	if opts.Consumables {
		preset.Consumables = modification.Consumables
	}
	if opts.ActivatedConsumables {
		preset.ActivatedConsumables = modification.ActivatedConsumables
	}
	if opts.Buffs {
		preset.Buffs = modification.Buffs
	}
	if opts.TargetDebuffs {
		preset.TargetDebuffs = modification.TargetDebuffs
	}
	if opts.Talents {
		preset.Talents = modification.Talents
	}
	if opts.Options {
		preset.Options = modification.Options
	}

	return preset
}
