package subscriber

import (
	"bytes"
	"compress/zlib"
	"io"
	"log/slog"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CmdrVasquess/eddnc"
	zmq "github.com/pebbe/zmq4"
)

var log = slog.Default()

func SetLog(l *slog.Logger) { log = l }

type EnqueueFunc func(c chan<- []byte, data []byte, scm eddnc.ScmID)

func Blocking(c chan<- []byte, data []byte, _ eddnc.ScmID) {
	c <- data
}

func Dropping(c chan<- []byte, data []byte, queue string) {
	select {
	case c <- data:
	default:
		log.Warn("dropping message from `queue`", "queue", queue)
	}
}

type DropStats struct{ Total, Dropped uint64 }

type DropWithStats map[eddnc.ScmID]*DropStats

func NewDropWithStats() DropWithStats { return make(DropWithStats) }

func (dws DropWithStats) Enqueue(c chan<- []byte, data []byte, queue eddnc.ScmID) {
	stats := dws[queue]
	if stats == nil {
		stats = new(DropStats)
		dws[queue] = stats
	}
	stats.Total++
	select {
	case c <- data:
	default:
		stats.Dropped++
		log.Warn("`drop` message of `total` from `queue`",
			"drop", stats.Dropped,
			"total", stats.Total,
			"queue", queue,
		)
	}
}

type S struct {
	Chan [eddnc.ScmNo]<-chan []byte

	rtuxm       int64
	chanNo      int
	relay       string
	connTimeout time.Duration
	recvTimeout time.Duration
	closing     int32
	enqueue     EnqueueFunc
}

const (
	DefaultRelay = "tcp://eddn.edcd.io:9500"
	GoodTimeout  = 6 * time.Second
)

type Config struct {
	Relay       string
	ConnTimeout time.Duration
	RecvTimeout time.Duration
	QCaps       [eddnc.ScmNo]int
	Enqueue     EnqueueFunc
}

// NoQ disables all channels, i.e. set the QCaps to -1.
func (cfg *Config) AllQCaps(cap int) *Config {
	for i := eddnc.ScmID(0); i < eddnc.ScmNo; i++ {
		cfg.QCaps[i] = cap
	}
	return cfg
}

// QCap sets the capactiy of the channel for schema ID q to cap.
func (cfg *Config) QCap(q eddnc.ScmID, cap int) *Config {
	cfg.QCaps[q] = cap
	return cfg
}

func New(cfg *Config) *S {
	res := &S{
		relay:       cfg.Relay,
		connTimeout: cfg.ConnTimeout,
		recvTimeout: cfg.RecvTimeout,
		enqueue:     cfg.Enqueue,
	}
	if res.recvTimeout == 0 {
		res.recvTimeout = -1
	}
	var chans [eddnc.ScmNo]chan<- []byte
	for i := eddnc.ScmID(0); i < eddnc.ScmNo; i++ {
		if cap := cfg.QCaps[i]; cap >= 0 {
			c := make(chan []byte, cap)
			chans[i] = c
			res.Chan[i] = c
			res.chanNo++
		}
	}
	if res.relay == "" {
		res.relay = DefaultRelay
	}
	if res.enqueue == nil {
		res.enqueue = Blocking
	}
	go res.loop(chans)
	return res
}

func (s *S) Return(rawEvent []byte) {
	bufPool.Put(rawEvent[:0])
}

func (s *S) UsedChannels() int { return s.chanNo }

func (s *S) TRecvUnixMilli() int64 { return atomic.LoadInt64(&s.rtuxm) }

func (s *S) Close() bool {
	return atomic.CompareAndSwapInt32(&s.closing, 0, 1)
}

func must(err error) {
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

var (
	bufPool  = sync.Pool{New: func() interface{} { return []byte{} }}
	scmMatch = regexp.MustCompile(`"\$schemaRef"[^:]*:\s*"([^"]+)"`)
)

func (s *S) loop(chans [eddnc.ScmNo]chan<- []byte) {
	zctx, err := zmq.NewContext()
	must(err)
	var subs *zmq.Socket
	defer func() {
		if subs != nil {
			subs.Close()
		}
	}()
	for {
		log.Debug("0MQ connecting to `relay`", "relay", s.relay)
		subs, err = zctx.NewSocket(zmq.SUB)
		must(err)
		must(subs.SetSubscribe(""))
		must(subs.SetConnectTimeout(s.connTimeout))
		must(subs.SetRcvtimeo(s.recvTimeout))
		must(subs.Connect(s.relay))
		for {
			if atomic.CompareAndSwapInt32(&s.closing, 1, -1) {
				for i, c := range chans {
					if c != nil {
						close(c)
						chans[i] = nil
					}
				}
				return
			}
			msg, err := subs.RecvBytes(0)
			if err != nil { // TODO Only break on EAGAIN?
				log.Error(err.Error())
				break
			}
			atomic.StoreInt64(&s.rtuxm, time.Now().UnixMilli())
			zrd, err := zlib.NewReader(bytes.NewReader(msg))
			if err != nil {
				log.Error(err.Error())
				continue
			}
			txt := bytes.NewBuffer(bufPool.Get().([]byte))
			io.Copy(txt, &io.LimitedReader{R: zrd, N: 5 * 1024 * 1024}) // Avoid ZIP bombing
			zrd.Close()
			line := txt.Bytes()
			var scm string
			if m := scmMatch.FindSubmatch(line); m == nil {
				log.Error("no $schemaRef in `message`", "message", string(line))
				continue
			} else {
				scm = string(m[1])
			}
			if scmid, ok := eddnc.ScmMap[scm]; ok {
				if c := chans[scmid]; c != nil {
					s.enqueue(c, line, scmid)
				}
			} else {
				bufPool.Put(line)
				if !strings.HasSuffix(scm, "/test") {
					log.Error("unknown `schema`", "schema", scm)
				}
			}
		}
		log.Debug("close 0MQ socket")
		subs.Close()
		subs = nil
	}
}
