package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"git.fractalqb.de/fractalqb/ggja"
)

func jFail(err error) {
	log.Fatal(err)
}

func rdJson(filename string) *ggja.Obj {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	obj := make(map[string]interface{})
	err = dec.Decode(&obj)
	if err != nil {
		log.Fatal(err)
	}
	return &ggja.Obj{Bare: obj, OnError: jFail}
}

func edcHasCommodity(edcCds *ggja.Arr, cdt string) string {
	for _, i := range edcCds.Bare {
		edc := ggja.Obj{Bare: i.(ggja.GenObj), OnError: edcCds.OnError}
		inm := edc.Str("locName", "")
		if inm == cdt {
			return edc.MStr("name")
		}
	}
	return ""
}

func makeMaps(catMap, cmtMap map[string]string, edce, jsmk *ggja.Obj) {
	edccds := edce.MObj("market").MArr("commodities")
	mkitms := jsmk.MArr("Items")
	for _, tm := range mkitms.Bare {
		mitm := ggja.Obj{Bare: tm.(ggja.GenObj), OnError: jsmk.OnError}
		mNmLoc := mitm.MStr("Name_Localised")
		edc := edcHasCommodity(edccds, mNmLoc)
		if len(edc) == 0 {
			fmt.Fprintf(os.Stderr, "cannot map market item '%s' to edc commodity\n", mNmLoc)
		} else {
			mNm := mitm.MStr("Name")
			fmt.Printf("M: \"%s\" \"%s\" (%s)\n", mNm, edc, mitm.Str("Category", "-"))
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
