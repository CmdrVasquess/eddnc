package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CmdrVasquess/eddnc"
)

var eddn = eddnc.Upload{
	Vaildate: true,
	TestUrl:  true,
	DryRun:   true,
}

func readJournal(file string) (good, bad int) {
	rd, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()
	line := 0
	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		line++
		entry := make(map[string]interface{})
		err = json.Unmarshal(scn.Bytes(), &entry)
		if err != nil {
			log.Fatal(err)
		}
		msg := eddnc.NewMessage(eddnc.Ts(time.Now()))
		err = eddnc.SetJournalJ(msg, entry, "TESTSYS", 4711, 1, 2, 3, false)
		if err != nil {
			continue
		}
		err = eddn.Send(eddnc.Sjournal, msg)
		if err != nil {
			bad++
			log.Printf("%s:%d:%s", file, line, err)
			log.Println("MESSAGE >>>>>>>>>>>>>>>>>>>>>")
			var m strings.Builder
			enc := json.NewEncoder(&m)
			enc.Encode(msg)
			log.Println(m.String())
			log.Println("MESSAGE <<<<<<<<<<<<<<<<<<<<<")
		} else {
			good++
		}
	}
	return good, bad
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		good, bad := readJournal(arg)
		fmt.Printf("%s: good=%d / bad=%d\n", arg, good, bad)
	}
}
