package eddn

import (
	"encoding/json"
	"fmt"
	"strings"

	"git.fractalqb.de/fractalqb/ggja"
)

var cmdtFromJMarket = map[string]string{
	"meanPrice":     "MeanPrice",
	"buyPrice":      "BuyPrice",
	"stock":         "Stock",
	"stockBracket":  "StockBracket",
	"sellPrice":     "SellPrice",
	"demand":        "Demand",
	"demandBracket": "DemandBracket",
}

func cmdtConvert(jMkt ggja.GenObj) (res ggja.GenObj, err error) {
	defer func() {
		if x := recover(); x != nil {
			res = nil
			j, _ := json.Marshal(jMkt)
			err = fmt.Errorf("%s: %s", err, string(j))
		}
	}()
	name := (&ggja.Obj{Bare: jMkt}).MStr("Name") // heavy just for error handling?
	res = make(ggja.GenObj)
	if strings.HasPrefix(name, "$") {
		name = name[1:]
	}
	if strings.HasSuffix(name, ";") {
		name = name[:len(name)-1]
	}
	if strings.HasSuffix(name, "_name") {
		name = name[:len(name)-5]
	}
	res["name"] = name
	for edcNm, mktNm := range cmdtFromJMarket {
		if tmp, ok := jMkt[mktNm]; ok {
			res[edcNm] = tmp
		}
	}
	return res, err
}

func SetCommoditiesJ(msg map[string]interface{}, journal map[string]interface{}) error {
	if tmp, ok := journal["StarSystem"]; ok {
		msg["systemName"] = tmp
	} else {
		return fmt.Errorf("missing system name in commodities data: %s", journal)
	}
	if tmp, ok := journal["StationName"]; ok {
		msg["stationName"] = tmp
	} else {
		return fmt.Errorf("missing station name in commodities data: %s", journal)
	}
	if tmp, ok := journal["MarketID"]; ok {
		msg["marketId"] = tmp
	}
	var items []interface{}
	if tmp, ok := journal["Items"]; ok {
		itmls := tmp.([]interface{})
		items = make([]interface{}, 0, len(itmls))
		for _, src := range itmls {
			si := src.(map[string]interface{})
			di, err := cmdtConvert(si)
			if err != nil {
				return err
			} else if di != nil {
				items = append(items, di)
			}
		}
	}
	msg["commodities"] = items
	return nil
}
