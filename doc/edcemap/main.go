package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"git.fractalqb.de/fractalqb/daq"
)

func rdJson(filename string) daq.DictAny {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	obj := make(daq.DictAny)
	err = dec.Decode(&obj)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func edcHasCommodity(edcCds daq.SliceAny, cdt string) string {
	for _, i := range edcCds {
		var edc daq.DictAny = i.(map[string]any)
		inm := edc.StringOr("locName", "")
		if inm == cdt {
			return daq.Must(edc.AsString("name"))
		}
	}
	return ""
}

func makeMaps(catMap, cmtMap map[string]string, edce, jsmk daq.DictAny) {
	edccds := daq.Must(daq.AsSliceAny(daq.Get(edce, "market", "commodities")))
	mkitms := daq.Must(jsmk.AsSliceAny("Items"))
	for _, tm := range mkitms {
		var mitm daq.DictAny = tm.(map[string]any)
		mNmLoc := daq.Must(mitm.AsString("Name_Localised"))
		edc := edcHasCommodity(edccds, mNmLoc)
		if len(edc) == 0 {
			fmt.Fprintf(os.Stderr, "cannot map market item '%s' to edc commodity\n", mNmLoc)
		} else {
			mNm := daq.Must(mitm.AsString("Name"))
			fmt.Printf("M: \"%s\" \"%s\" (%s)\n", mNm, edc, mitm.StringOr("Category", "-"))
		}
	}
}

func main() {
	edce := rdJson(os.Args[1])
	jsmk := rdJson(os.Args[2])
	catMap := make(map[string]string)
	cmtMap := make(map[string]string)
	makeMaps(catMap, cmtMap, edce, jsmk)
}
