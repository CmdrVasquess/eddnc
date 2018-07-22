package eddn

import (
	"fmt"
)

func SetShipyardJ(msg map[string]interface{}, journal map[string]interface{}) error {
	if tmp, ok := journal["StarSystem"]; ok {
		msg["systemName"] = tmp
	} else {
		return fmt.Errorf("missing system name in shipyard data: %s", journal)
	}
	if tmp, ok := journal["StationName"]; ok {
		msg["stationName"] = tmp
	} else {
		return fmt.Errorf("missing station name in shipyard data: %s", journal)
	}
	tmp, ok := journal["PriceList"]
	if !ok {
		return fmt.Errorf("missing price list in shipyard data: $s", journal)
	}
	list := tmp.([]interface{})
	items := make([]string, 0, len(list))
	for _, tmp := range list {
		item := tmp.(map[string]interface{})
		shty := item["ShipType"].(string)
		items = append(items, shty)
	}
	msg["ships"] = items
	return nil
}
