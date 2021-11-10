#=IF(CX33+CY33+CZ33+DB33+DC33+DA33=0,"None",IF(S32>=U32, "Auto",INDEX($CX$30:$DC$30,0, MATCH(Max(CX33:DC33),CX33:DC33,0))))
def prio_spell():
    if steady + auto + Multi + raptor + melee == 0:
        return "None"
    else if mana_spent >= mana_pool:
        return "Auto"
    else:
        return name_of(MAX(steady, auto, multi, raptor, melee))

#=CD32/((MAX(AR33,0)*(CB32*$CX$31)+CF32))
def steady_weight():
    return steady_damage / (steady_cooldown * speed * 1.2 + steady_cast_time)


#=CA32/(MAX(AQ33,0)*CB32*$CY$31+CC32)
def auto_weight():
    return auto_dmg / (auto_cd * speed * 1.4 + auto_cast_time)

