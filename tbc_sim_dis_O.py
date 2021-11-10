#IF(BA33<=0,11.6,20),IF(BG33<=0,10,20)))

if mana_asc == 18 and proc_rdy == 20:
    return 18
else:
    return MIN(
        if TBW <= 0:
            return 1
        else:
            return 20
        ,
        if racial <= 0:
            return 2
        else:
            return 20
        ,
        if trinket_1_is_proc:
            return 20 # proc_rdy
        else if trinket_1 <= 0:
            return 3
        else:
            return 20
        ,
        if trinket_2_is_proc:
            return 20
        else if T2 >= 0:
            return 6
        else:
            return 20
        ,
        if rapid_fire <= 0:
            return 4
        else:
            return 20
        ,
        if ability_prio == "Auto":
            return 18
        else if ability_prio == "Steady":
            return 19
        else if ability_prio == "Multi-Shot":
            return 14
        else if ability_prio == "raptor":
            return 13
        else if ability_prio == "Melee":
            return 15
        else if ability_prio == "Arcane Shot":
            return 17
        else:
            return 20
        ,
        PROC_RDY (Can be either 20, 5 or 6), which is None, Quick Shots and DST
        ,
        if bl <= 0:
            return 8
        else:
            return 20
        ,
        if pot <= 0:
            return 9
        else:
            return 20
        ,
        if rune <= 0:
            return 11.5
        else:
            return 20
        ,
        if readiness <= 0:
            return 11.6
        else:
            return 20
        ,
        if drum <= 0:
            return 10
        else:
            return 20
    )