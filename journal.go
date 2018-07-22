package eddn

import (
	"fmt"
	"strings"
)

func attDrop(name string) bool {
	return strings.HasSuffix(name, "_Localised")
}

func setExcept(msg, journal map[string]interface{}, overwrite bool, drop ...string) {
NEXT_KEY:
	for k, v := range journal {
		if attDrop(k) {
			continue
		}
		for _, d := range drop {
			if d == k {
				continue NEXT_KEY
			}
		}
		jsonObjSet(msg, k, v, overwrite)
	}
}

func SetJournal(msg, journal map[string]interface{}, sys string, x, y, z float64, overwrite bool) error {
	msg["StarSystem"] = sys
	msg["StarPos"] = []float64{x, y, z}
	evtNm := journal["event"].(string)
	switch evtNm {
	case "FSDJump":
		setExcept(msg, journal, overwrite, "BoostUsed", "FuelLevel", "FuelUsed", "JumpDist")
	case "Scan":
		setExcept(msg, journal, overwrite)
	case "Docked":
		setExcept(msg, journal, overwrite, "CockpitBreach")
	case "Location":
		setExcept(msg, journal, overwrite, "Latitude", "Longitude")
	default:
		return fmt.Errorf("EDDN does not accept journal event '%s'", evtNm)
	}
	return nil
}
