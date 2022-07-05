package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
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
	generate()
}

type schemaInfo struct {
	ID string
}

func generate() {
	fmt.Printf(`// generated with go run gen/main.go
package eddnc

import _ "embed"

const ScmNo = %d

//go:generate stringer -type ScmID
const (
`, len(schemaFiles))
	for i := range schemaFiles {
		fmt.Printf("\tS%s ScmID = %d\n", schemaName[i], i)
	}
	fmt.Print(`)

var ScmURLs = []string{
`)
	uris := make([]string, len(schemaFiles))
	for i, f := range schemaFiles {
		schema, err := ioutil.ReadFile(filepath.Join(config.Schemas, f))
		if err != nil {
			panic(err)
		}
		var sinfo schemaInfo
		json.Unmarshal(schema, &sinfo)
		uris[i] = strings.TrimRight(sinfo.ID, "#")
		fmt.Printf("\t\"%s\",\n", uris[i])
	}
	fmt.Print(`}

var ScmMap = map[string]ScmID{
`)
	for i := range uris {
		fmt.Printf("\t\"%s\": S%s,\n", uris[i], schemaName[i])
	}
	fmt.Print(`}

var ScmDefs = []string{
`)
	for _, n := range schemaName {
		fmt.Printf("\t%sSchema,\n", n)
	}
	fmt.Print(`}

var (
`)
	for i, n := range schemaName {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf(`	//go:embed %s
	%sSchema string
`, filepath.Join("./schemas", schemaFiles[i]), n)
	}
	fmt.Println(")")
}
