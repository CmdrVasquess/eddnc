package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"git.fractalqb.de/fractalqb/catch"
	"github.com/CmdrVasquess/eddnc"
	"github.com/CmdrVasquess/eddnc/subscriber"
	jsoniter "github.com/json-iterator/go"
)

var (
	qLengths = 32
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	files    = make(map[string]*os.File)

	fSplice string
)

func main() {
	flag.IntVar(&qLengths, "q", qLengths, "Set length of internal event queues")
	flag.StringVar(&fSplice, "splice", fSplice, "Splice event to files in dir")
	flag.Parse()

	subs := subscriber.New((&subscriber.Config{
		ConnTimeout: subscriber.GoodTimeout,
		RecvTimeout: 30 * time.Second,
		Enqueue:     subscriber.Dropping,
	}).AllQCaps(qLengths)) // => eventLoop() must read all channels to avoid blocking
	// Blocking could also be avoided by setting a non-blocking Enqueue function with
	// Config (see Blocking, Dropping).
	go eventLoop(subs)
	// Be polite and clean up…
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("^C: shutting down...")
	subs.Close()
	for n, f := range files {
		if f != nil {
			log.Println("close ", n)
			f.Close()
		}
	}
	time.Sleep(time.Second) // just to see eventLoop exit
	log.Println("Bye")
}

func eventLoop(subs *subscriber.S) {
	wr := os.Stdout
	openChans := subs.UsedChannels()
	dump := func(scm eddnc.ScmID, msg []byte, ok bool) {
		if len(msg) > 0 {
			wr.Write(msg)
			fmt.Fprintln(wr)
			splice(scm, msg)
			subs.Return(msg)
		}
		if !ok {
			openChans--
		}
	}
	for openChans > 0 {
		select {
		case b, ok := <-subs.Chan[eddnc.Sapproachsettlement]:
			dump(eddnc.Sapproachsettlement, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sblackmarket]:
			dump(eddnc.Sblackmarket, b, ok)
		case b, ok := <-subs.Chan[eddnc.Scodexentry]:
			dump(eddnc.Scodexentry, b, ok)
		case b, ok := <-subs.Chan[eddnc.Scommodity]:
			dump(eddnc.Scommodity, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sdockingdenied]:
			dump(eddnc.Sdockingdenied, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sdockinggranted]:
			dump(eddnc.Sdockinggranted, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfcmaterials_capi]:
			dump(eddnc.Sfcmaterials_capi, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfcmaterials_journal]:
			dump(eddnc.Sfcmaterials_journal, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssallbodiesfound]:
			dump(eddnc.Sfssallbodiesfound, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssbodysignals]:
			dump(eddnc.Sfssbodysignals, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfssdiscoveryscan]:
			dump(eddnc.Sfssdiscoveryscan, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sfsssignaldiscovered]:
			dump(eddnc.Sfsssignaldiscovered, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sjournal]:
			dump(eddnc.Sjournal, b, ok)
		case b, ok := <-subs.Chan[eddnc.Snavbeaconscan]:
			dump(eddnc.Snavbeaconscan, b, ok)
		case b, ok := <-subs.Chan[eddnc.Snavroute]:
			dump(eddnc.Snavroute, b, ok)
		case b, ok := <-subs.Chan[eddnc.Soutfitting]:
			dump(eddnc.Soutfitting, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sscanbarycentre]:
			dump(eddnc.Sscanbarycentre, b, ok)
		case b, ok := <-subs.Chan[eddnc.Sshipyard]:
			dump(eddnc.Sshipyard, b, ok)
		}
	}
	log.Println("exit event loop")
}

func splice(scm eddnc.ScmID, msg []byte) {
	if fSplice == "" {
		return
	}
	message := json.Get(msg, "message").ToString()
	if scm == eddnc.Sjournal {
		event := json.Get([]byte(message), "event").ToString()
		n := filepath.Join(fSplice, scm.String())
		os.MkdirAll(n, 0777)
		f := getFile(filepath.Join(n, event+".json"))
		f.WriteString(message)
		fmt.Fprintln(f)
	} else {
		n := filepath.Join(fSplice, scm.String()+".json")
		f := getFile(n)
		f.WriteString(message)
		fmt.Fprintln(f)
	}
}

func getFile(n string) *os.File {
	f := files[n]
	if f == nil {
		f = catch.MustRet(os.Create(n))
		files[n] = f
	}
	return f
}
