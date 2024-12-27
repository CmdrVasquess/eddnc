package eddnc

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"git.fractalqb.de/fractalqb/eloc"
	"git.fractalqb.de/fractalqb/eloc/must"
	zmq "github.com/pebbe/zmq4"
)

//go:generate go run ./gen schemas.go
//go:generate gofmt -s -w schemas.go

type state = int32

const (
	stClean state = iota
	stStarting
	stRunning
	stShutdown
	stClosing
	stDone
)

const (
	TheRelay    = "tcp://eddn.edcd.io:9500"
	ConnTimeout = 6 * time.Second
	RecvTimeout = 2 * time.Second
	NoTimout    = time.Duration(-1)
)

func NqDrop(q chan<- []byte, b []byte) {
	select {
	case q <- b:
	default:
	}
}

// One-time-use; create a new after Stop()
type Subscriber struct {
	Relay       string
	ConnTimeout time.Duration
	RecvTimeout time.Duration
	Nq          func(chan<- []byte, []byte)
	LogInfo     func(format string, args ...any)
	LogError    func(format string, args ...any)

	zmqctx *zmq.Context
	sock   *zmq.Socket
	trecv  int64

	state state
}

func (zc *Subscriber) LastRecv() time.Time {
	uxm := atomic.LoadInt64(&zc.trecv)
	return time.UnixMilli(uxm)
}

func (zc *Subscriber) Run(q chan<- []byte) (err error) {
	if zc.ConnTimeout < 0 {
		zc.ConnTimeout = 0
	} else if zc.ConnTimeout == 0 {
		zc.ConnTimeout = ConnTimeout
	}
	if zc.RecvTimeout < 0 {
		zc.RecvTimeout = 0
	} else if zc.RecvTimeout == 0 {
		zc.RecvTimeout = RecvTimeout
	}

	if !atomic.CompareAndSwapInt32(&zc.state, stClean, stStarting) {
		return eloc.New("cannot run client")
	}

	defer func() {
		p := recover()

		atomic.StoreInt32(&zc.state, stClosing)
		if zc.sock != nil {
			if err := zc.sock.Close(); err != nil {
				zc.lerrf("0mq %p socket close: %s", zc, err)
			}
			zc.sock = nil
		}
		if zc.zmqctx != nil {
			if err := zc.zmqctx.Term(); err != nil {
				zc.lerrf("0mq %p context terminate: %s", zc, err)
			}
			zc.zmqctx = nil
		}
		close(q)

		if p != nil {
			switch p := p.(type) {
			case error:
				err = p
			case string:
				err = errors.New(p)
			default:
				err = fmt.Errorf("panic: %+v", p)
			}
		}
		if err != nil {
			zc.lerrf("0mq %p #%d stopped with error: %s", zc, len(q), err)
		} else {
			zc.linff("0mq %p #%d stopped", zc, len(q))
		}
		atomic.StoreInt32(&zc.state, stDone)
	}()

	zc.linff("0mq %p create context", zc)
	zc.zmqctx = must.Ret(zmq.NewContext())

	zc.linff("0mq %p create socket", zc)
	zc.sock = must.Ret(zc.zmqctx.NewSocket(zmq.SUB))

	must.Do(zc.sock.SetSubscribe(""))
	must.Do(zc.sock.SetConnectTimeout(zc.ConnTimeout))
	must.Do(zc.sock.SetRcvtimeo(zc.RecvTimeout))

	zc.linff("0mq %p to %s", zc, zc.Relay)
	must.Do(zc.sock.Connect(zc.Relay))

	nq := func(q chan<- []byte, b []byte) { q <- b }
	if zc.Nq != nil {
		nq = zc.Nq
	}

	atomic.StoreInt32(&zc.state, stRunning)
	for atomic.LoadInt32(&zc.state) == stRunning {
		msg, err := zc.sock.RecvBytes(0)
		if err != nil {
			if errors.Is(err, zmq.Errno(11)) {
				continue
			}
			return err
		}
		nq(q, msg)
		atomic.StoreInt64(&zc.trecv, time.Now().UnixMilli())
	}

	return nil
}

func (zc *Subscriber) Stop() {
	for !atomic.CompareAndSwapInt32(&zc.state, stRunning, stShutdown) {
		if atomic.LoadInt32(&zc.state) > stRunning {
			return
		}
		time.Sleep(time.Millisecond)
	}
}

func (zc *Subscriber) linff(fmt string, args ...any) {
	if zc.LogInfo != nil {
		zc.LogInfo(fmt, args...)
	}
}

func (zc *Subscriber) lerrf(fmt string, args ...any) {
	if zc.LogError != nil {
		zc.LogError(fmt, args...)
	}
}
