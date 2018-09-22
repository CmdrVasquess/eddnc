package eddn

import (
	"fmt"
	"strings"

	"git.fractalqb.de/fractalqb/nmconv"
)

var outfitNmc = nmconv.Conversion{
	Norm:   nmconv.Unsep("_"),
	Xform:  nmconv.PerSegment(nmconv.CapWord),
	Denorm: nmconv.Sep("_"),
}

func SetOutfittingJ(msg map[string]interface{}, journal map[string]interface{}) error {
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
	tmp, ok := journal["Items"]
	if !ok {
		return fmt.Errorf("missing items in outfitting data: %s", journal)
	}
	list := tmp.([]interface{})
	items := make([]string, 0, len(list))
	for _, tmp := range list {
		item := tmp.(map[string]interface{})
		name := item["Name"].(string)
		switch {
		case strings.HasPrefix(name, "int_") && name != "int_planetapproachsuite":
			fallthrough
		case strings.HasPrefix(name, "hpt_"):
			fallthrough
		case strings.Index(name, "_armour_") >= 0:
			items = append(items, outfitNmc.Convert(name))
		}
	}
	msg["modules"] = items
	return nil
}
