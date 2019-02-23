package eddn

import (
	"fmt"
	"strings"
)

func attDrop(name string) bool {
	return strings.HasSuffix(name, "_Localised")
}

var (
	blFactions = map[string]interface{}{"MyReputation": true}

	blJFsdJump = map[string]interface{}{
		"BoostUsed": true,
		"FuelLevel": true,
		"FuelUsed":  true,
		"JumpDist":  true,
		"Factions":  blFactions,
	}

	blJDocked = map[string]interface{}{
		"CockpitBreach": true,
	}

	blJLocation = map[string]interface{}{
		"Latitude":  true,
		"Longitude": true,
		"Factions":  blFactions,
	}
)

func SetJournalJ(
	msg, journal map[string]interface{},
	sys string,
	sysAddr uint64,
	x, y, z float64,
	overwrite bool,
) error {
	msg["StarSystem"] = sys
	msg["SystemAddress"] = sysAddr
	msg["StarPos"] = []float64{x, y, z}
	evtNm := journal["event"].(string)
	switch evtNm {
	case "FSDJump":
		setBl(msg, journal, overwrite, blJFsdJump)
	case "Scan":
		setBl(msg, journal, overwrite, blEmpty)
	case "Docked":
		setBl(msg, journal, overwrite, blJDocked)
	case "Location":
		setBl(msg, journal, overwrite, blJLocation)
	default:
		return fmt.Errorf("EDDN does not accept journal event '%s'", evtNm)
	}
	return nil
}
