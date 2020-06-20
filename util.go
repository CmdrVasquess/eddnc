package eddnc

import (
	"fmt"
)

func jsonObjSet(dst map[string]interface{}, key string, val interface{}, overwrite bool) {
	if overwrite {
		dst[key] = val
	} else if _, ok := dst[key]; !ok {
		dst[key] = val
	}
}

func jsonObjSubmap(dst map[string]interface{}, key string) (res map[string]interface{}) {
	elm, ok := dst[key]
	if ok {
		res = elm.(map[string]interface{}) // FIXME should be checked
	} else {
		res = make(map[string]interface{})
		dst[key] = res
	}
	return res
}

var blEmpty = map[string]interface{}{}

func setBl(msg, journal map[string]interface{}, overwrite bool, blacklist map[string]interface{}) {
	for k, v := range journal {
		bl, ok := blacklist[k]
		if !ok {
			jsonObjSet(msg, k, v, overwrite)
		} else if subbl, ok := bl.(map[string]interface{}); ok {
			switch subj := v.(type) {
			case map[string]interface{}:
				subm := jsonObjSubmap(msg, k)
				setBl(subm, subj, overwrite, subbl)
			case []interface{}:
				msl := make([]interface{}, 0, len(subj))
				for _, je := range subj {
					me := make(map[string]interface{})
					setBl(me, je.(map[string]interface{}), overwrite, subbl)
					msl = append(msl, me)
				}
				msg[k] = msl
			default:
				panic(fmt.Errorf("setBl cannot copy '%v'", v))
			}
		}
	}
}
