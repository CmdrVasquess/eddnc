package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"git.fractalqb.de/fractalqb/eloc/must"
)

var (
	config = struct {
		Schemas string
	}{
		Schemas: "doc/EDDN/schemas",
	}

	schemaFilePattern = regexp.MustCompile(`(.+)-v.+\.json`)
	schemaFiles       []string
	schemaName        []string
)

func main() {
	ls, _ := os.ReadDir(config.Schemas)
	for _, e := range ls {
		match := schemaFilePattern.FindStringSubmatch(e.Name())
		if match == nil {
			continue
		}
		schemaFiles = append(schemaFiles, e.Name())
		schemaName = append(schemaName, match[1])
		sort.Strings(schemaFiles)
	}
	if len(os.Args) == 1 {
		generate(os.Stdout)
	} else {
		w := must.Ret(os.Create(os.Args[1]))
		defer w.Close()
		generate(w)
	}
}

type schemaInfo struct {
	ID string
}

func generate(w io.Writer) {
	uris := make([]string, len(schemaFiles))
	for i, f := range schemaFiles {
		schema, err := os.ReadFile(filepath.Join(config.Schemas, f))
		if err != nil {
			panic(err)
		}
		var sinfo schemaInfo
		json.Unmarshal(schema, &sinfo)
		uris[i] = strings.TrimRight(sinfo.ID, "#")
	}
	packed, uripos := packstr(uris...)

	fmt.Fprintf(w, `// generated with go run gen/main.go
package eddnc

const ScmNo = %d

type ScmID int

func (sid ScmID) Info() ScmInfo { return ScmInfos[sid] }

//go%s stringer -type ScmID
const (
`, len(schemaFiles), ":generate")
	for i := range schemaFiles {
		fmt.Fprintf(w, "\tS%s ScmID = %d\n", schemaName[i], i)
	}
	fmt.Fprint(w, `)

type ScmInfo struct {
	Ref     string
	Topic   string
	Version int
}

var ScmInfos = []ScmInfo{
`)
	for i, uri := range uris {
		pos := uripos[i]
		top, v := refsplit(uri)
		fmt.Fprintf(w, "\t{allScms[%d:%d], allScms[%d:%d], %d}, // %s\n",
			pos.start, pos.end,
			pos.start, pos.end-top,
			v,
			uri,
		)
	}
	fmt.Fprint(w, `}

var ScmMap = map[string]ScmID{
`)
	for i := range uris {
		pos := uripos[i]
		fmt.Fprintf(w, "\tallScms[%d:%d]: S%s,\n",
			pos.start, pos.end,
			schemaName[i],
		)
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintf(w, "\nconst allScms = \"%s\"\n", packed)

	/* Do not embed schemas

	var ScmDefs = []string{
	`)
		for _, n := range schemaName {
			fmt.Fprintf(w, "\t%sSchema,\n", n)
		}
		fmt.Fprint(w, `}

	var (
	`)
		   	for i, n := range schemaName {
		   		if i > 0 {
		   			fmt.Fprintln(w)
		   		}
		   		fmt.Fprintf(w, `	//go:embed %s
		   	%sSchema string

		   `, filepath.Join("./schemas", schemaFiles[i]), n)

		   	}
		   	fmt.Fprintln(w, ")")
	*/
}
