package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/CmdrVasquess/goEDDNc/down"
)

func eventLoop(subs *down.Subscriber) {
	enc := json.NewEncoder(os.Stdout)
LOOP:
	for {
		select {
		case b, ok := <-subs.Blackmarket:
			if ok {
				enc.Encode(b)
			} else {
				break LOOP
			}
		case c, ok := <-subs.Commodity:
			if ok {
				enc.Encode(c)
			} else {
				break LOOP
			}
		case j, ok := <-subs.Journal:
			if ok {
				enc.Encode(j)
			} else {
				break LOOP
			}
		case o, ok := <-subs.Outfitting:
			if ok {
				enc.Encode(o)
			} else {
				break LOOP
			}
		case s, ok := <-subs.Shipyard:
			if ok {
				enc.Encode(s)
			} else {
				break LOOP
			}
		}
	}
	log.Println("exit event loop")
}

func main() {
	subs := down.New(down.Config{Timeout: down.GoodTimeout})
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
