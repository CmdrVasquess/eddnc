package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"git.fractalqb.de/fractalqb/c4hgol"
	"github.com/CmdrVasquess/eddnc"
	"github.com/CmdrVasquess/eddnc/subscriber"
)

var qLengths = 16

func eventLoop(subs *subscriber.S) {
	wr := os.Stdout
	openChans := subs.UsedChannels()
	dump := func(msg []byte, ok bool) {
		if len(msg) > 0 {
			wr.Write(msg)
			fmt.Fprintln(wr)
			subs.Return(msg)
		}
		if !ok {
			openChans--
		}
	}
	for openChans > 0 {
		select {
		case b, ok := <-subs.Chan[eddnc.Sblackmarket]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Scodexentry]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Scommodity]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssdiscoveryscan]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sjournal]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Snavbeaconscan]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Snavroute]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Soutfitting]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sscanbarycentre]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sshipyard]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sapproachsettlement]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfcmaterials_capi]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfcmaterials_journal]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssallbodiesfound]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssbodysignals]:
			dump(b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfsssignaldiscovered]:
			dump(b, ok)
		}
	}
	log.Println("exit event loop")
}

func main() {
	flag.IntVar(&qLengths, "q", qLengths, "Set length of internal event queues")
	flag.Parse()

	c4hgol.Configure(subscriber.LogCfg, "t", true)

	subs := subscriber.New((&subscriber.Config{
		ConnTimeout: subscriber.GoodTimeout,
		RecvTimeout: 30 * time.Second,
	}).AllQCaps(qLengths)) // => eventLoop() must read all channels to avoid blocking
	// Blocking could also be avoided by setting a non-blocking Enqueue function with
	// Config (see Blocking, Dropping).
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
