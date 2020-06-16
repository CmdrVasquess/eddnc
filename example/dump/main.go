package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/CmdrVasquess/goEDDNc/down"
)

func eventLoop(subs *down.Subscriber) {
	enc := json.NewEncoder(os.Stdout)
	openChans := subs.UsedChannels()
	for openChans > 0 {
		select {
		case b, ok := <-subs.Blackmarket:
			if ok {
				enc.Encode(b)
			} else {
				openChans--
			}
		case c, ok := <-subs.Commodity:
			if ok {
				enc.Encode(c)
			} else {
				openChans--
			}
		case j, ok := <-subs.Journal:
			if ok {
				enc.Encode(j)
			} else {
				openChans--
			}
		case o, ok := <-subs.Outfitting:
			if ok {
				enc.Encode(o)
			} else {
				openChans--
			}
		case s, ok := <-subs.Shipyard:
			if ok {
				enc.Encode(s)
			} else {
				openChans--
			}
		}
	}
	log.Println("exit event loop")
}

func writeMemStats() {
	wr, _ := os.Create("eddn-dump.memstats")
	defer wr.Close()
	ticks := time.NewTicker(10 * time.Second)
	var mstat runtime.MemStats
	for {
		<-ticks.C
		runtime.ReadMemStats(&mstat)
		fmt.Fprintf(wr, "%+v\n", &mstat)
	}
}

func main() {
	subs := down.New(down.Config{Timeout: down.GoodTimeout})
	go writeMemStats()
	go eventLoop(subs)
	// Be polite and clean up…
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("^C: shutting down...")
	subs.Close()
	time.Sleep(time.Second) // just to see eventLoop exit
	log.Println("Bye")
}
