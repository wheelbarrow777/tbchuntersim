# =BS40*BU40*(1+BT40/HasteRatingRatio/100)

def steady_cast_time():
    return base_haste * haste_bonus * (1 + haste_bonus_rating * 100 / HasteRatingRatio)


#=IF(AO31>0,1.4,1)*IF(AE31>0,1.15,1)*IF(AF31>0,1.3,1)*IF(AA31>0,1+$I$20,1)
def haste_bonus():
    ratio = 1
    if rapid_fire_buff_exp > 0:
        ratio = ratio * 1.4
    else:
        ratio = ratio * 1
    
    if quick_shots_buff_exp > 0:
        ratio = ratio * 1.15
    else:
        ratio = ratio * 1

    if bloodlust_buff_exp > 0:
        ratio = ratio * 1.3
    else:
        ratio = ratio * 1

    if racial_buff_exp > 0:
        ratio = ratio * (1 + racial_bonus_haste) # This is not the case for orc
    else:
        ratio = ratio * 1
    
    return ratio

#=IF(AG34>0,$J$25,0)+IF(AH34>0,$J$26,0)
def haste_bonus_rating():
    r = haste_rating

    if trinket_1_cd > 0:
        r =  r + trinket_1_haste

        
    if trinket_2_cd > 0:
        r =  r + trinket_2_haste
    
    if haste_pot_cd > 0:
        r = r + haste_pot_haste
    
    if drum_cd > 0:
        r = r + drum_haste