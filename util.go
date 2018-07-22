package eddn

func jsonObjSet(dst map[string]interface{}, key string, val interface{}, overwrite bool) {
	if overwrite {
		dst[key] = val
	} else if _, ok := dst[key]; !ok {
		dst[key] = val
	}
}
