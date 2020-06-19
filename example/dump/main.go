package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/CmdrVasquess/goEDDNc/subscriber"
)

func eventLoop(subs *subscriber.S) {
	wr := os.Stdout
	openChans := subs.UsedChannels()
	for openChans > 0 {
		select {
		case b, ok := <-subs.Blackmarket:
			if ok {
				wr.Write(b)
				subs.Return(b)
			} else {
				openChans--
			}
		case c, ok := <-subs.Commodity:
			if ok {
				wr.Write(c)
				subs.Return(c)
			} else {
				openChans--
			}
		case j, ok := <-subs.Journal:
			if ok {
				wr.Write(j)
				subs.Return(j)
			} else {
				openChans--
			}
		case o, ok := <-subs.Outfitting:
			if ok {
				wr.Write(o)
				subs.Return(o)
			} else {
				openChans--
			}
		case s, ok := <-subs.Shipyard:
			if ok {
				wr.Write(s)
				subs.Return(s)
			} else {
				openChans--
			}
		}
		fmt.Fprintln(wr)
	}
	log.Println("exit event loop")
}

func main() {
	subs := subscriber.New(subscriber.Config{
		Timeout: subscriber.GoodTimeout,
	})
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
