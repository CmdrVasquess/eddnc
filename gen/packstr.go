package main

import (
	"strconv"
	"strings"

	"git.fractalqb.de/fractalqb/eloc/must"
)

type strpos struct{ start, end int }

func packstr(strs ...string) (string, []strpos) {
	var (
		all strings.Builder
		ps  = make([]strpos, len(strs))
	)
	for i, s := range strs {
		p := strpos{start: all.Len()}
		all.WriteString(s)
		p.end = all.Len()
		ps[i] = p
	}
	return all.String(), ps
}

func refsplit(ref string) (end, v int) {
	end = strings.LastIndexByte(ref, '/')
	v = must.Ret(strconv.Atoi(string(ref[end+1:])))
	end = len(ref) - end - 1
	return end, v
}
