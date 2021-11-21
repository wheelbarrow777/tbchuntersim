package itemdb

import (
	"fmt"
	"strings"
	eq "tbchuntersim/equipment"
)

func getItem(name string, m interface{}) eq.ArmorItem {
	name = strings.ToLower(name)
	if doesItemExist(name, m) {
		if itemMap, ok := m.(map[string]eq.ArmorItem); ok {
			return itemMap[name]
		} else {
			return eq.ArmorItem{}
		}
	} else {
		panic(fmt.Sprintf("%s does not exist in the db", name))
	}
}
