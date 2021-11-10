#=IF(VLOOKUP(O33,CLReferenceTable,14,FALSE)="YES",X33+1.5,IF(X33<Z32,Z32,0))
def gcd():
    if ability_has_gcd:
        return ability_start_time + 1.5
    else if ability_start_time < gcd_prev:
        return gcd_prev
    else:
        return 0

#=IF(IF(Q33="Raptor",MAX(AX33,0)+$AG$22,IF(Q33="Melee",MAX(AY33,0)+$AG$22,IF(Q33="Arcane Shot",MAX(AZ33,0)+$AG$22,0))))))
def start_time():
    var = 0
    if action == "Auto":
        var = MAX(auto_remaining_cd, 0)
    else if action == "Steady":
        var = MAX(steady_remaining_cd, 0) + latency
    else if action == "Multi-Shot":
        var = MAX(multi_remaining_cd, 0) + latency
    else if action == "raptor":
        return MAX(raptor_remaining_cd, 0) + latency
    else if action == "melee":
        return MAX(melee_remaining_cd, 0) + latency
    else if action == "arcane shot":
        return MAX(arcane_shot_cd, 0) + latency
    return stop_time_prev + var

#=IF(MIN(MAX(AQ33,0),MAX(AR33,0),MAX(AS33,0),0),0)))))))
def stop_time():
    var = 0
    if action == "auto":
        var = autoshot_cast_time
    else if action == "steady shot":
        var = steady_cast_time
    else if action == "mutli-shot":
        var = multi-shot_cast_time
    else if action == "raptor":
        var = melee_weave_time
    else if action == "melee":
        var = melee_weave_time
    else if action == "none":
        var = MIN(auto_remaining_cd, steady_remaining_cd, multi_remaining_cd)
    
    return start_time + var

    

#=MAX(AR32-(Y32-Y31),0))+IF(Q32="Multi-Shot",Z32-Y32,0)+IF(Q32="Arcane Shot",Z32-Y32,0))
def steady_cd():
    if prev_action == "Readiness":
        return 1
    else if prev_action == "Steady":
        return gcd_prev - stop_time_prev
    else:
        var = 0
        if prev_action == "Multi-Shot":
            var = gcd_prev - stop_time_prev
        else if prev_action == "Arcane Shot":
            var = gcd_prev - stop_time_prev
        return steady_cd - (stop_time_prev - stop_time_prev_prev) + var


#=IF(Q32="Auto",CB32-CC32,AQ32-(Y32-Y31))
def auto_cd():
    if prev-action == "auto":
        return auto_speed - auto_cast_time
    else:
        return auto_remaining_cd - (stop_time - stop_time_prev) # auto_remaining_cd - time_elapsed