package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/CmdrVasquess/eddnc"
)

func main() {
	zc := eddnc.Subscriber{
		Relay:       eddnc.TheRelay,
		ConnTimeout: 5 * time.Second,
		RecvTimeout: time.Second,
		Nq:          eddnc.NqDrop,
		LogInfo:     log.Printf,
		LogError:    log.Printf,
	}

	var shutdown sync.WaitGroup

	msgs := make(chan []byte, 128)
	shutdown.Add(1)
	go func() {
		var line []byte
		for msg := range msgs {
			var err error
			line, err = eddnc.Decode(msg, 0, line[:0]) // production s/w should set a limit
			if err != nil {
				log.Print(err)
			} else {
				fmt.Println(string(line))
			}
		}
		log.Print("exit message loop")
		shutdown.Done()
	}()

	shutdown.Add(1)
	go func() {
		zc.Run(msgs)
		shutdown.Done()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("shutdown…")
	zc.Stop()
	shutdown.Wait()
	log.Println("o7")
}
